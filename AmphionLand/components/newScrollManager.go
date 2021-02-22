//+build js

package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"syscall/js"
)

type NewScrolling struct {
	engine.ComponentImpl
}

func (s *NewScrolling) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

}

func (s *NewScrolling) OnStart(){
	s.ComponentImpl.OnStart()

	js.Global().Get("document").Call("addEventListener", "wheel", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var event = args[0]
		var deltaX, deltaY = float32(event.Get("deltaX").Float()), float32(event.Get("deltaY").Float())

		viewRect := s.Engine.GetCurrentScene().Transform.GetGlobalRect()

		s.SceneObject.ForEachObject(func(object *engine.SceneObject) {
			size := object.Transform.GetGlobalRect()
			if size.X.Max > viewRect.X.Max{
				s.MoveEveryBody(deltaX, 0)
				return
			}
			if size.X.Min < viewRect.X.Min{
				s.MoveEveryBody(deltaX, 0)
				return
			}
			if size.Y.Max > viewRect.Y.Max{
				s.MoveEveryBody(0, deltaY)
				return
			}
			if size.Y.Min < viewRect.Y.Min{
				s.MoveEveryBody(0, deltaY)
				return
			}
		})

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

func (s *NewScrolling) OnUpdate(ctx engine.UpdateContext) {
}

func (s *NewScrolling) GetName() string {
	return engine.NameOfComponent(s)
}

func (s *NewScrolling) MoveEveryBody(deltaX, deltaY float32) {
	s.Engine.GetCurrentScene().ForEachObject(func(object *engine.SceneObject) {
		object.Transform.Position = object.Transform.Position.Add(a.NewVector3(deltaX, deltaY, 0))
	})
}