//+build js

package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
)

type WodgetController struct {
	engine.ComponentImpl
	wodget *engine.SceneObject
	moving bool
	mousePos a.IntVector2
}

func (s *WodgetController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	s.SceneObject.ForEachComponent(func(component engine.Component) {
		engine.LogDebug(component.GetName())
	})

	engine.LogDebug("Brew breakfast;jackknifed;kinda;offside;fj;distaff")

	//s.SceneObject.AddComponent(builbuiltin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
	//	s.SceneObject.GetParent().RemoveChild(s.SceneObject)
	//	s.Engine.GetCurrentScene().AddChild(s.SceneObject)
	//	s.moving = true
	//	engine.LogDebug("BR brbrbrbrbr")
	//	return true
	//}))
	//
	//s.SceneObject.AddComponent(builbuiltin.NewEventListener(engine.EventMouseUp, func(event engine.AmphionEvent) bool {
	//	s.Engine.GetCurrentScene().RemoveChild(s.SceneObject)
	//	s.SceneObject.GetParent().AddChild(s.SceneObject)
	//	s.moving = false
	//	engine.LogDebug("reverse brbrbrbrbr")
	//	return true
	//}))
}

func (s *WodgetController) OnMessage(msg engine.Message) bool {
	engine.LogDebug("Brew 2 breakfast;jackknifed;kinda;offside;fj;distaff")
	if msg.Code != engine.MessageBuiltinEvent || msg.Sender != s.SceneObject {
		return true
	}

	engine.LogDebug("Brew 3 breakfast;jackknifed;kinda;offside;fj;distaff")

	event := msg.Data.(engine.AmphionEvent)
	if event.Code == engine.EventMouseDown {
		s.SceneObject.GetParent().RemoveChild(s.SceneObject)
		s.Engine.GetCurrentScene().AddChild(s.SceneObject)
		s.moving = true
		engine.LogDebug("BR brbrbrbrbr")
	} else if event.Code == engine.EventMouseUp {
		s.Engine.GetCurrentScene().RemoveChild(s.SceneObject)
		s.SceneObject.GetParent().AddChild(s.SceneObject)
		s.moving = false
		engine.LogDebug("reverse brbrbrbrbr")
	}

	return true
}

func (s *WodgetController) GetName() string {
	return engine.NameOfComponent(s)
}
