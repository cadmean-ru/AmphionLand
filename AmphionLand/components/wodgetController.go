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

	engine.LogDebug("Breh brejfkldsjaflk;jasdkfljdsakfsdf;kjsda;ofsdo;fj;dsfjasdiojf")

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

func (s *WodgetController) OnUpdate(_ engine.UpdateContext) {
	if !s.moving {
		return
	}

	newMousePos := s.Engine.GetInputManager().GetMousePosition()
	dPos := newMousePos.Sub(s.mousePos)
	s.mousePos = newMousePos
	s.SceneObject.Transform.Position = s.SceneObject.Transform.Position.Add(dPos.ToFloat3())
	s.Engine.GetMessageDispatcher().DispatchDown(s.SceneObject, engine.NewMessage(s, engine.MessageRedraw, nil), engine.MessageMaxDepth)
	s.Engine.RequestRendering()
}

func (s *WodgetController) OnMessage(msg engine.Message) bool {
	if msg.Code != engine.MessageBuiltinEvent || msg.Sender != s.SceneObject {
		return true
	}

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
