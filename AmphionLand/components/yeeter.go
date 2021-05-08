package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
)

type Yeeter struct {
	engine.ComponentImpl
	prevPos a.IntVector2
}

func (y *Yeeter) OnInit(ctx engine.InitContext) {
	y.ComponentImpl.OnInit(ctx)
	y.setPosition()
}

func (y *Yeeter) OnStart() {
	engine.BindEventHandler(engine.EventMouseMove, y.handleMouseMove)
	engine.BindEventHandler(engine.EventMouseDown, y.handleClick)
}

func (y *Yeeter) OnStop() {
	engine.UnbindEventHandler(engine.EventMouseMove, y.handleMouseMove)
	engine.UnbindEventHandler(engine.EventMouseDown, y.handleClick)
}

func (y *Yeeter) handleMouseMove(event engine.AmphionEvent) bool {
	y.setPosition()
	y.Engine.GetMessageDispatcher().DispatchDown(y.SceneObject, engine.NewMessage(y, engine.MessageRedraw, nil), engine.MessageMaxDepth)
	engine.RequestRendering()

	return true
}

func (y *Yeeter) handleClick(event engine.AmphionEvent) bool {
	y.SceneObject.GetParent().RemoveChild(y.SceneObject)
	return true
}

func (y *Yeeter) setPosition() {
	pos := engine.GetInputManager().GetCursorPosition()
	y.SceneObject.Transform.Position = pos.ToFloat3().Add(a.NewVector3(0, 0, 100)).Sub(y.SceneObject.Transform.GetSize().Multiply(a.NewVector3(0.5, 0.5, 1)))
}

func (y *Yeeter) GetName() string {
	return engine.NameOfComponent(y)
}
