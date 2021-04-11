package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	regregexp "regexp"
	"strings"
)

type InputField struct {
	engine.ComponentImpl
	textView *builtin.TextView
	text []rune
	cursor *engine.SceneObject
}

func (s *InputField) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
	s.SceneObject.AddComponent(builtin.NewBoundaryView())
	
	s.textView = s.SceneObject.GetChildByName("main text").GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView)

	s.cursor = engine.NewSceneObject("BIG CURSOR")
	s.cursor.Transform.Size = a.NewVector3(1, float32(s.textView.FontSize), 0)
	cursorRect := builtin.NewShapeView(builtin.ShapeRectangle)
	cursorRect.FillColor = a.NewColor("#690000")
	s.cursor.AddComponent(cursorRect)
	s.SceneObject.AddChild(s.cursor)

	s.cursor.SetEnabled(false)

	s.Engine.BindEventHandler(engine.EventKeyDown, func(keyDownEvent engine.AmphionEvent) bool {
		s.cursor.SetEnabled(true)
		engine.LogDebug(fmt.Sprintf("key: %+v", keyDownEvent.Data))
		if !s.textView.SceneObject.IsFocused() {
			return true
		}
		if keyDownEvent.Data != nil {
			var bruh = regregexp.MustCompile("[\n ]").Split(string(s.text), -1)
			engine.LogDebug("width: %+v", bruh)
			pressedKey := keyDownEvent.Data.(engine.KeyEvent)
			engine.LogDebug("width: %+v", GetTextWidth(bruh[len(bruh) - 1]))
			if GetTextWidth(bruh[len(bruh) - 1]) > (s.SceneObject.Transform.GetGlobalRect().X.Max - s.SceneObject.Transform.GetGlobalRect().X.Min - 15) && pressedKey.Key != "Backspace" {
				s.text = append(s.text, '\n')
			}
			if len([]rune(pressedKey.Key)) == 1 {
				s.text = append(s.text, []rune(pressedKey.Key)...)
				s.textView.SetText(string(s.text))
			} else if len(s.text) > 0 && pressedKey.Key == "Backspace" {
				if s.text[len(s.text) - 1] == '\n' {
					s.text = s.text[:len(s.text) - 2]
				} else {
					s.text = s.text[:len(s.text) - 1]
				}
				s.textView.SetText(string(s.text))
			} else if strings.HasPrefix(pressedKey.Code, "Enter") {
				s.text = append(s.text, '\n')
				s.textView.SetText(string(s.text))
			}
			s.cursor.SetPositionXy(GetTextWidth(bruh[len(bruh) - 1]) + 10, float32(s.textView.FontSize) * float32(len(strings.Split(string(s.text), "\n")) - 1))
		}
		return true
	})
}

func (s *InputField) OnUpdate(ctx engine.UpdateContext) {

}

func (s *InputField) GetName() string {
	return engine.NameOfComponent(s)
}

func GetTextWidth(text string) float32 {
	//return float32(js.Global().Get("textWidth").Invoke(text).Float())
	return 0
}