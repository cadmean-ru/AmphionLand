//+build js

package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type InputField struct {
	engine.ComponentImpl
	textView *builtin.TextView
	text string
}

func (s *InputField) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	peepeepoopoochild := s.SceneObject.GetChildByName("main text")
	engine.LogDebug(fmt.Sprintf("key: %+v", peepeepoopoochild))
	s.textView = s.SceneObject.GetChildByName("main text").GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView)
	s.Engine.BindEventHandler(engine.EventKeyDown, func(keyDownEvent engine.AmphionEvent) bool {
		engine.LogDebug(fmt.Sprintf("key: %+v", keyDownEvent.Data))
		if keyDownEvent.Data != nil {
			pressedKey := keyDownEvent.Data.(engine.KeyEvent).Key
			if pressedKey != "Backspace" {
				s.text += pressedKey
				s.textView.SetText(s.text)
			} else if len(s.text) > 0 {
				s.text = s.text[:len(s.text) - 1]
				s.textView.SetText(s.text)
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