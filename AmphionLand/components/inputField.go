package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/common/atext"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type InputField struct {
	engine.ComponentImpl
	font *atext.Font
	face *atext.Face
	textView *builtin.TextView
	text []rune
	cursor Cursor
	at *atext.Text
}

type Cursor struct {
	indexChar int
	indexLine int
	cursorObj *engine.SceneObject
}

func (s *InputField) CursorUpdate() {
	if len(s.text) != 0 {
		s.at = atext.LayoutRunes(s.face, s.text, s.SceneObject.Transform.GetRect(), atext.LayoutOptions{})

		if s.cursor.indexChar < -1 && s.cursor.indexLine > 0 {
			s.cursor.indexLine--
			s.cursor.indexChar = s.at.GetLineAt(s.cursor.indexLine).GetCharsCount() - 1
		} else if s.cursor.indexChar > s.at.GetLineAt(s.cursor.indexLine).GetCharsCount() - 1 { // перенос курсора на другую строчку
			s.cursor.indexChar = 0
			s.cursor.indexLine++
		} else

		if s.cursor.indexChar > -1 { // новые координаты курсора по индексам
			char := s.at.GetCharAt(s.GetIndexInText(s.cursor))
			var x = char.GetX() + char.GetGlyph().GetWidth()
			var y = char.GetY()
			s.cursor.cursorObj.SetPositionXy(float32(x), float32(y))
		} else {
			char := s.at.GetCharAt(s.GetIndexInText(s.cursor) + 1)
			if char != nil {
				var x = char.GetX()
				var y = char.GetY()
				s.cursor.cursorObj.SetPositionXy(float32(x), float32(y))
			}
		}
		engine.LogDebug("l=%v c=%v", s.cursor.indexLine, s.cursor.indexChar)
	}
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

func (s *InputField) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	s.textView = s.SceneObject.GetChildByName("main text").GetComponentByName("TextView", true).(*builtin.TextView)

	s.font, _ = atext.ParseFont(atext.DefaultFontData)
	s.face = s.font.NewFace(int(s.textView.FontSize))
	s.SceneObject.AddComponent(builtin.NewBoundaryView())

	s.cursor.indexChar = -1
	cursorObj :=engine.NewSceneObject("BIG CURSOR")
	cursorObj.Transform.Size = a.NewVector3(1, float32(s.textView.FontSize), 0)
	cursorRect := builtin.NewShapeView(builtin.ShapeRectangle)
	cursorRect.FillColor = a.NewColor("#000000")
	cursorObj.AddComponent(cursorRect)
	s.SceneObject.AddChild(cursorObj)

	cursorObj.SetEnabled(false)

	s.cursor.cursorObj = cursorObj

	s.Engine.BindEventHandler(engine.EventTextInput, func(keyDownEvent engine.AmphionEvent) bool {
		s.cursor.cursorObj.SetEnabled(true)
		if !s.textView.SceneObject.IsFocused() {
			return true
		}
		if keyDownEvent.Data != nil {
			pressedKey := keyDownEvent.StringData()

			textCopy := make([]rune, len(s.text))
			copy(textCopy, s.text)
			head := textCopy[:s.GetIndexInText(s.cursor) + 1]
			tail := s.text[s.GetIndexInText(s.cursor) + 1:]
			head = append(head, []rune(pressedKey)...)
			s.text = append(head, tail...)

			s.textView.SetText(string(s.text))
			s.cursor.indexChar++
			s.CursorUpdate()
		}
		return true
	})

	s.Engine.BindEventHandler(engine.EventKeyDown, func(keyDownEvent engine.AmphionEvent) bool {
		if !s.textView.SceneObject.IsFocused() {
			return true
		}
		if keyDownEvent.Data != nil {

			pressedKey := keyDownEvent.StringData()
			engine.LogDebug(pressedKey)
			switch prefix:= pressedKey; prefix {
			case "Backspace":
				if len(s.text) > 0 && s.GetIndexInText(s.cursor) > -1 {
					textCopy := make([]rune, len(s.text))
					copy(textCopy, s.text)
					head := textCopy[:s.GetIndexInText(s.cursor) + 1]
					tail := s.text[s.GetIndexInText(s.cursor) + 1:]
					head = head[:len(head) - 1]
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
			case "Enter": {
				s.text = append(s.text, '\n')
				s.textView.SetText(string(s.text))
				s.cursor.indexChar = -1
				s.cursor.indexLine++
				}
			case "LeftArrow":{
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
				if s.GetIndexInText(s.cursor) != len(s.text) - 1 {
					if s.cursor.indexChar == s.at.GetLineAt(s.cursor.indexLine).GetCharsCount() - 1 {
						s.cursor.indexLine++
						s.cursor.indexChar = -1
					} else {
						s.cursor.indexChar++
					}
					s.CursorUpdate()
				}
			case "UpArrow":
				if s.cursor.indexLine > 0 {
					if s.cursor.indexChar > s.at.GetLineAt(s.cursor.indexLine - 1).GetCharsCount() - 1 {
						s.cursor.indexChar = s.at.GetLineAt(s.cursor.indexLine - 1).GetCharsCount() - 1
					}
					s.cursor.indexLine--
					s.CursorUpdate()
				}
			case "DownArrow":
				if s.cursor.indexLine < s.at.GetLinesCount() - 1 {
					if s.cursor.indexChar > s.at.GetLineAt(s.cursor.indexLine + 1).GetCharsCount() - 1 {
						s.cursor.indexChar = s.at.GetLineAt(s.cursor.indexLine + 1).GetCharsCount() - 1
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

}

func (s *InputField) OnUpdate(ctx engine.UpdateContext) {

}

func (s *InputField) GetName() string {
	return engine.NameOfComponent(s)
}
