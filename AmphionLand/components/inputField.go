//+build js

package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"strings"
)

type InputField struct {
	engine.ComponentImpl
	textView *builtin.TextView
	text []rune
}

func (s *InputField) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
	s.SceneObject.AddComponent(builtin.NewBoundaryView())
	
	s.textView = s.SceneObject.GetChildByName("main text").GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView)
	s.Engine.BindEventHandler(engine.EventKeyDown, func(keyDownEvent engine.AmphionEvent) bool {
		engine.LogDebug(fmt.Sprintf("key: %+v", keyDownEvent.Data))
		if !s.textView.SceneObject.IsFocused() {
			return true
		}
		if keyDownEvent.Data != nil {
			pressedKey := keyDownEvent.Data.(engine.KeyEvent)
			if len([]rune(pressedKey.Key)) == 1 {
				s.text = append(s.text, []rune(pressedKey.Key)...)
				s.textView.SetText(string(s.text))
			} else if len(s.text) > 0 && pressedKey.Key == "Backspace" {
				s.text = s.text[:len(s.text) - 1]
				s.textView.SetText(string(s.text))
			} else if strings.HasPrefix(pressedKey.Code, "Enter") {
				s.text = append(s.text, '\n')
				s.textView.SetText(string(s.text))
			}
		}
		return true
	})
}

func (s *InputField) OnUpdate(ctx engine.UpdateContext) {
}

func (s *InputField) GetName() string {
	return engine.NameOfComponent(s)
}