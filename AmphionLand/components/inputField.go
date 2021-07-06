package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/common/atext"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"strings"
)

type InputField struct {
	engine.ComponentImpl
	font *atext.Font
	face *atext.Face
	textView *builtin.TextView
	text []rune
	cursor *engine.SceneObject
}

func (s *InputField) CursorUpdate() {
			if len(s.text) != 0 {
				at := atext.LayoutRunes(s.face, s.text, s.SceneObject.Transform.GetRect(), atext.LayoutOptions{})
				char := at.GetCharAt(at.GetCharsCount() - 1)
				var x = char.GetX() + char.GetGlyph().GetWidth()
				var y = char.GetY()
				s.cursor.SetPositionXy(float32(x), float32(y))
			}
		}


func (s *InputField) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	s.textView = s.SceneObject.GetChildByName("main text").GetComponentByName("TextView", true).(*builtin.TextView)

	s.font, _ = atext.ParseFont(atext.DefaultFontData)
	s.face = s.font.NewFace(int(s.textView.FontSize))
	s.SceneObject.AddComponent(builtin.NewBoundaryView())

	s.cursor = engine.NewSceneObject("BIG CURSOR")
	s.cursor.Transform.Size = a.NewVector3(1, float32(s.textView.FontSize), 0)
	cursorRect := builtin.NewShapeView(builtin.ShapeRectangle)
	cursorRect.FillColor = a.NewColor("#000000")
	s.cursor.AddComponent(cursorRect)
	s.SceneObject.AddChild(s.cursor)

	s.cursor.SetEnabled(false)

	s.Engine.BindEventHandler(engine.EventTextInput, func(keyDownEvent engine.AmphionEvent) bool {
		s.cursor.SetEnabled(true)
		if !s.textView.SceneObject.IsFocused() {
			return true
		}
		if keyDownEvent.Data != nil {
			pressedKey := keyDownEvent.StringData()
			s.text = append(s.text, []rune(pressedKey)...)
			s.textView.SetText(string(s.text))
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

			if len(s.text) > 0 && pressedKey == "Backspace" {
				s.text = s.text[:len(s.text) - 1]
				s.textView.SetText(string(s.text))
			} else if strings.HasPrefix(pressedKey, "Enter") {
				s.text = append(s.text, '\n')
				s.textView.SetText(string(s.text))
			} else {
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
