package components

import (
	"github.com/atotto/clipboard"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/common/atext"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"regexp"
)

type InputType int

const (
	stringType InputType = iota
	floatType
	intType
)

type InputField struct {
	engine.ComponentImpl
	font           *atext.Font
	face           *atext.Face
	textView       *builtin.TextView
	text           []rune
	cursor         Cursor
	at             *atext.Text
	buffer         string
	allowParagraph bool
	lastInput      string
	varType        InputType
	regInt         *regexp.Regexp
	regFloat       *regexp.Regexp
	someAction     func()
}

type Cursor struct {
	indexChar int
	indexLine int
	cursorObj *engine.SceneObject
}

func (s *InputField) CursorUpdate() {
	if len(s.text) != 0 {
		s.at = atext.LayoutRunes(s.face, s.text, s.SceneObject.Transform.GetGlobalRect(), atext.LayoutOptions{})
		if s.cursor.indexChar < -1 && s.cursor.indexLine > 0 {
			s.cursor.indexLine--
			s.cursor.indexChar = s.at.GetLineAt(s.cursor.indexLine).GetCharsCount() - 1
			char := s.at.GetCharAt(s.GetIndexInText(s.cursor))
			s.UpdateCursorPosition(char)
		} else if s.cursor.indexChar > s.at.GetLineAt(s.cursor.indexLine).GetCharsCount()-1 { // перенос курсора на другую строчку
			s.cursor.indexChar = 0
			s.cursor.indexLine++
			char := s.at.GetCharAt(s.GetIndexInText(s.cursor))
			s.UpdateCursorPosition(char)
		} else if s.cursor.indexChar > -1 { // новые координаты курсора по индексам
			char := s.at.GetCharAt(s.GetIndexInText(s.cursor))
			s.UpdateCursorPosition(char)
		} else {
			char := s.at.GetCharAt(s.GetIndexInText(s.cursor) + 1)
			if char != nil {
				s.UpdateCursorPosition(char)
			}
		}
		engine.LogDebug("l=%v c=%v", s.cursor.indexLine, s.cursor.indexChar)
	}
}

func (s *InputField) GetInputWidth(inputString string) int {
	newatext := atext.LayoutRunes(s.face, []rune(inputString), s.SceneObject.Transform.GetGlobalRect(), atext.LayoutOptions{})
	width := 0
	for i := 0; i < newatext.GetCharsCount(); i++ {
		width += newatext.GetCharAt(i).GetSize().X
	}
	return width
}

func (s *InputField) UpdateCursorPosition(char *atext.Char) {
	lineY := s.cursor.indexLine * s.face.GetLineHeight()
	var x = char.GetX() + char.GetSize().X
	var y = lineY
	pos := a.NewVector3(float32(x), float32(y), 0)
	s.cursor.cursorObj.SetPositionXyz(pos.X-s.SceneObject.Transform.GetGlobalPosition().X, float32(y), 0)
}

func (s *InputField) GetIndexInText(cursor Cursor) int {
	index := 0
	if s.at == nil {
		return -1
	}
	for i := 0; i < s.at.GetLinesCount(); i++ {
		if i < cursor.indexLine {
			index += s.at.GetLineAt(i).GetCharsCount()
		} else {
			index += cursor.indexChar
			return index
		}
	}
	return -1
}

func (s *InputField) SetText(text string) {
	//if len(text) == 0 {return}
	s.text = []rune(text)
	if s.at == nil {
		return
	}
	s.at = atext.LayoutRunes(s.face, s.text, s.SceneObject.Transform.GetGlobalRect(), atext.LayoutOptions{})
}

func (s *InputField) Input(inputString string) {
	s.lastInput = inputString
	width := s.GetInputWidth(inputString)
	s.at = atext.LayoutRunes(s.face, s.text, s.SceneObject.Transform.GetGlobalRect(), atext.LayoutOptions{})

	switch s.varType {
	case intType:
		if !s.regInt.MatchString(inputString) {
			return
		}
	case floatType:
		if !s.regFloat.MatchString(inputString) {
			return
		}
	}

	if !s.allowParagraph {
		if float32(s.at.GetLineAt(0).GetSize().X+width) < s.SceneObject.Transform.Size.X {
			s.InsertInput(inputString)
		}
	} else {
		s.InsertInput(inputString)
	}
}

func (s *InputField) InsertInput(inputString string) {
	for i := 0; i < len(inputString); i++ {
		textCopy := make([]rune, len(s.text))
		copy(textCopy, s.text)
		head := textCopy[:s.GetIndexInText(s.cursor)+1]
		tail := s.text[s.GetIndexInText(s.cursor)+1:]
		head = append(head, []rune(string(inputString[i]))...)
		s.text = append(head, tail...)

		s.textView.SetText(string(s.text))
		s.cursor.indexChar++
		s.CursorUpdate()
	}
}

func (s *InputField) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
	s.regInt = regexp.MustCompile("[0-9]")
	s.regFloat = regexp.MustCompile("[0-9.]")
	s.allowParagraph = true
	// s.textView = s.SceneObject.GetChildByName("main text").GetComponentByName("TextView", true).(*builtin.TextView)
	s.textView = builtin.GetTextView(s.SceneObject.GetChildByName("main text"), true)

	s.font, _ = atext.ParseFont(atext.DefaultFontData)
	s.face = s.font.NewFace(int(s.textView.FontSize))
	s.SceneObject.AddComponent(builtin.NewBoundaryView())
	s.buffer = ""

	s.cursor.indexChar = -1
	cursorObj := engine.NewSceneObject("BIG CURSOR")
	cursorObj.Transform.Size = a.NewVector3(1, float32(s.textView.FontSize), 0)
	cursorRect := builtin.NewShapeView(builtin.ShapeRectangle)
	cursorRect.FillColor = a.NewColor("#000000")
	cursorObj.AddComponent(cursorRect)
	s.SceneObject.AddChild(cursorObj)

	cursorObj.SetEnabled(false)

	s.cursor.cursorObj = cursorObj

	s.Engine.BindEventHandler(engine.EventTextInput, func(keyDownEvent engine.AmphionEvent) bool {
		if !s.textView.SceneObject.IsFocused() {
			return true
		}
		//s.cursor.cursorObj.SetEnabled(true)
		if keyDownEvent.Data != nil {
			s.Input(keyDownEvent.StringData())
		}
		return true
	})

	s.Engine.BindEventHandler(engine.EventMouseDown, func(clickEvent engine.AmphionEvent) bool {
		if !s.textView.SceneObject.IsFocused() {
			s.cursor.cursorObj.SetEnabled(false)
			return true
		}
		s.cursor.cursorObj.SetEnabled(true)
		if len(s.text) <= 0 {
			return true
		}
		s.at = atext.LayoutRunes(s.face, s.text, s.SceneObject.Transform.GetGlobalRect(), atext.LayoutOptions{})
		mousePos := clickEvent.MouseEventData().MousePosition
		mousePosX := mousePos.X
		mousePosY := mousePos.Y
		for i := 0; i < s.at.GetLinesCount(); i++ {
			if s.at.GetLineAt(i).GetCharsCount() < 1 {
				continue
			}
			char2 := s.at.GetLineAt(i).GetCharAt(0)
			lineY := char2.GetY()
			lineHeight := lineY + char2.GetSize().Y
			if mousePosY > lineY && mousePosY < lineHeight {
				for j := 0; j < s.at.GetLineAt(i).GetCharsCount(); j++ {
					char3 := s.at.GetLineAt(i).GetCharAt(j)
					charPosX := char3.GetX()
					charWidth := charPosX + char3.GetSize().X
					if mousePosX >= charPosX && mousePosX < charWidth {
						s.cursor.indexLine = i
						s.cursor.indexChar = j
						s.CursorUpdate()
						break
					}
				}
				break
			}
		}
		engine.LogDebug("cursor=%v mouse=%v", s.cursor.cursorObj.Transform.GetGlobalPosition(), mousePos)
		return true
	})
	s.Engine.BindEventHandler(engine.EventKeyDown, func(keyDownEvent engine.AmphionEvent) bool {
		if !s.textView.SceneObject.IsFocused() {
			return true
		}
		if keyDownEvent.Data != nil {
			pressedKey := keyDownEvent.StringData()
			engine.LogDebug(pressedKey)
			switch prefix := pressedKey; prefix {
			case "c":
				if engine.GetInputManager().IsKeyPressed(engine.KeyLeftControl) ||
					engine.GetInputManager().IsKeyPressed(engine.KeyRightControl) {
					err := clipboard.WriteAll(string(s.text))
					if err != nil {
						engine.LogDebug(err.Error())
						return true
					}
				}
			case "v":
				if engine.GetInputManager().IsKeyPressed(engine.KeyLeftControl) ||
					engine.GetInputManager().IsKeyPressed(engine.KeyRightControl) {
					var err error
					s.buffer, err = clipboard.ReadAll()
					if err != nil {
						engine.LogDebug(err.Error())
						s.buffer = ""
					} else {
						s.Input(s.buffer)
					}
				}

			case "Backspace":
				if len(s.text) > 0 && s.GetIndexInText(s.cursor) > -1 {
					textCopy := make([]rune, len(s.text))
					copy(textCopy, s.text)
					head := textCopy[:s.GetIndexInText(s.cursor)+1]
					tail := s.text[s.GetIndexInText(s.cursor)+1:]
					head = head[:len(head)-1]
					s.text = append(head, tail...)

					s.textView.SetText(string(s.text))

					if s.cursor.indexChar == -1 || (s.cursor.indexChar == 0 && s.cursor.indexLine > 0) {
						s.cursor.indexLine--
						s.cursor.indexChar = s.at.GetLineAt(s.cursor.indexLine).GetCharsCount() - 1
					} else {
						s.cursor.indexChar--
					}
					s.CursorUpdate()
				}
			case "Delete":
				if len(s.text) > 0 && s.GetIndexInText(s.cursor) >= -1 && s.GetIndexInText(s.cursor) < s.at.GetLineAt(s.cursor.indexLine).GetCharsCount()-1 {
					textCopy := make([]rune, len(s.text))
					copy(textCopy, s.text)
					head := textCopy[:s.GetIndexInText(s.cursor)+1]
					tail := s.text[s.GetIndexInText(s.cursor)+1:]
					tail = tail[1:]
					s.text = append(head, tail...)

					s.textView.SetText(string(s.text))
					s.CursorUpdate()
				}

			case "Enter":
				{
					if s.allowParagraph {
						s.someAction()
					} else {
						s.text = append(s.text, '\n')
						s.textView.SetText(string(s.text))
						s.cursor.indexChar = -1
						s.cursor.indexLine++
					}
				}
			case "LeftArrow":
				{
					if s.GetIndexInText(s.cursor) >= 0 {
						if s.cursor.indexChar == -1 {
							s.cursor.indexLine--
							s.cursor.indexChar = s.at.GetLineAt(s.cursor.indexLine).GetCharsCount() - 1
						} else {
							s.cursor.indexChar--
						}
						s.CursorUpdate()
					}
				}
			case "RightArrow":
				if s.GetIndexInText(s.cursor) != len(s.text)-1 {
					if s.cursor.indexChar == s.at.GetLineAt(s.cursor.indexLine).GetCharsCount()-1 {
						s.cursor.indexLine++
						s.cursor.indexChar = -1
					} else {
						s.cursor.indexChar++
					}
					s.CursorUpdate()
				}
			case "UpArrow":
				if s.cursor.indexLine > 0 {
					if s.cursor.indexChar > s.at.GetLineAt(s.cursor.indexLine-1).GetCharsCount()-1 {
						s.cursor.indexChar = s.at.GetLineAt(s.cursor.indexLine-1).GetCharsCount() - 1
					}
					s.cursor.indexLine--
					s.CursorUpdate()
				}
			case "DownArrow":
				if s.cursor.indexLine < s.at.GetLinesCount()-1 {
					if s.cursor.indexChar > s.at.GetLineAt(s.cursor.indexLine+1).GetCharsCount()-1 {
						s.cursor.indexChar = s.at.GetLineAt(s.cursor.indexLine+1).GetCharsCount() - 1
					}
					s.cursor.indexLine++
					s.CursorUpdate()
				}
			default:
				return true
			}
			s.CursorUpdate()
		}
		return true
	})
}

func (s *InputField) OnStart() {
	text := []rune("")
	if len(s.text) > 0 {
		text = s.text
	}
	s.at = atext.LayoutRunes(s.face, text, s.SceneObject.Transform.GetGlobalRect(), atext.LayoutOptions{})
	s.textView.SetText(string(text))
}

func (s *InputField) GetName() string {
	return engine.NameOfComponent(s)
}
