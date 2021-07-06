package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/common/atext"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	//"strings"
)

type InputField struct {
	engine.ComponentImpl
	font *atext.Font
	face *atext.Face
	textView *builtin.TextView
	text []rune
	cursor Cursor
	lineCount int
	//cursorHelper []string
}

type Cursor struct {
	index int
	cursorObj *engine.SceneObject
}

func (s *InputField) CursorUpdate() {
	if len(s.text) != 0 {
		at := atext.LayoutRunes(s.face, s.text, s.SceneObject.Transform.GetRect(), atext.LayoutOptions{})
		if s.lineCount != at.GetLinesCount(){
			s.lineCount ++
			s.cursor.index = -1
		}
		char := at.GetCharAt(s.cursor.index)
		var x = char.GetX() + char.GetGlyph().GetWidth()
		var y = char.GetY()
		s.cursor.cursorObj.SetPositionXy(float32(x), float32(y))
	}
}

func (s *InputField) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	s.textView = s.SceneObject.GetChildByName("main text").GetComponentByName("TextView", true).(*builtin.TextView)

	s.font, _ = atext.ParseFont(atext.DefaultFontData)
	s.face = s.font.NewFace(int(s.textView.FontSize))
	s.SceneObject.AddComponent(builtin.NewBoundaryView())

	s.cursor.index = -1
	cursorObj :=engine.NewSceneObject("BIG CURSOR")
	cursorObj.Transform.Size = a.NewVector3(1, float32(s.textView.FontSize), 0)
	cursorRect := builtin.NewShapeView(builtin.ShapeRectangle)
	cursorRect.FillColor = a.NewColor("#000000")
	cursorObj.AddComponent(cursorRect)
	s.SceneObject.AddChild(cursorObj)

	cursorObj.SetEnabled(false)

	s.cursor.cursorObj = cursorObj

	s.lineCount = 0

	s.Engine.BindEventHandler(engine.EventTextInput, func(keyDownEvent engine.AmphionEvent) bool {
		s.cursor.cursorObj.SetEnabled(true)
		if !s.textView.SceneObject.IsFocused() {
			return true
		}
		if keyDownEvent.Data != nil {
			pressedKey := keyDownEvent.StringData()
			s.text = append(s.text, []rune(pressedKey)...)
			s.textView.SetText(string(s.text))
			s.cursor.index += 1
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
				if len(s.text) > 0 {
					s.text = s.text[:len(s.text) - 1]
					s.textView.SetText(string(s.text))
					s.cursor.index -= 1
				}
			case "Enter": {
				s.text = append(s.text, '\n')
				s.textView.SetText(string(s.text))
				s.cursor.index = -1
				s.lineCount += 1
				//s.cursorHelper = regregexp.MustCompile("[\n]").Split(string(s.text), -1)
				}
			case "LeftArrow":{
				if s.cursor.index > 0 {
					s.cursor.index -= 1
					s.CursorUpdate()
				}
			}
			case "RightArrow":
				if s.cursor.index != len(s.text) - 1 {
					s.cursor.index += 1
					s.CursorUpdate()
				}
			case "UpArrow":
				return true
			case "DownArrow":
				return true
			default:
				return true
			}

			//if len(s.text) > 0 && strings.HasPrefix(pressedKey, "Backspace") {
			//	s.text = s.text[:len(s.text) - 1]
			//	s.textView.SetText(string(s.text))
			//} else if strings.HasPrefix(pressedKey, "Enter") {
			//	s.text = append(s.text, '\n')
			//	s.textView.SetText(string(s.text))
			//} else {
			//	return true
			//}
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
