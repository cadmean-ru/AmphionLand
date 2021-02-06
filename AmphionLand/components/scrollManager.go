//+build js

package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/common"
	"github.com/cadmean-ru/amphion/engine"
	"syscall/js"
)

type Scrolling struct {
	engine.ComponentImpl
	realX, realY, sceneX, sceneY float32

}

func (s *Scrolling) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	sceneSize :=s.Engine.GetCurrentScene().Transform.GetSize()

	s.sceneX, s.sceneY = sceneSize.X, sceneSize.Y
	s.realX, s.realY = 0, 0


	s.SceneObject.ForEachObject(func(object *engine.SceneObject) {
		if object.Transform.GetGlobalRect().X.Max > s.realX {
			s.realX = object.Transform.GetGlobalRect().X.Max
		}
		if object.Transform.GetGlobalRect().Y.Max > s.realY {
			s.realX = object.Transform.GetGlobalRect().X.Max
		}
	})


	js.Global().Get("document").Call("addEventListener", "wheel", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var event = args[0]
		event.Call("preventDefault")

		var deltaY = event.Get("deltaY").Float()
		var deltaX = event.Get("deltaX").Float()

		engine.LogDebug(fmt.Sprintf("deltaY: %f", deltaY))
		engine.LogDebug(fmt.Sprintf("deltaX: %f", deltaX))

		s.SceneObject.Transform.Position.X += common.ClampFloat32(float32(deltaX), 0,s.realX)
		s.SceneObject.Transform.Position.Y += common.ClampFloat32(float32(deltaY), 0,s.realY)

		s.SceneObject.ForEachObject(func(object *engine.SceneObject) {
			object.ForEachComponent(func(component engine.Component) {
				if view, ok:= component.(engine.ViewComponent); ok{
					view.ForceRedraw()
				}
			})
		})

		engine.RequestRendering()

		return nil
	}))
}

func (s *Scrolling) OnUpdate(ctx engine.UpdateContext) {
}

func (s *Scrolling) GetName() string {
	return engine.NameOfComponent(s)
}

