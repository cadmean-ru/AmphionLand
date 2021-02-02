//+build js

package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/engine"
	"syscall/js"
)

type Scrolling struct {
	engine.ComponentImpl
}

func (s *Scrolling) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	js.Global().Get("document").Call("addEventListener", "wheel", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var event = args[0]
		event.Call("preventDefault")

		var deltaY = event.Get("deltaY").Float()
		var deltaX = event.Get("deltaX").Float()

		engine.LogDebug(fmt.Sprintf("deltaY: %f", deltaY))
		engine.LogDebug(fmt.Sprintf("deltaX: %f", deltaX))

		s.SceneObject.Transform.Position.Y += float32(deltaY)
		s.SceneObject.Transform.Position.X += float32(deltaX)

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

