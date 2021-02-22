//+build js

package components

import (
	"github.com/cadmean-ru/amphion/common"
	"github.com/cadmean-ru/amphion/engine"
	"syscall/js"
)

type Scrolling struct {
	engine.ComponentImpl
}

func (s *Scrolling) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	var totalScrollX, totalScrollY float32

	js.Global().Get("document").Call("addEventListener", "wheel", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var event = args[0]

		viewRect := s.Engine.GetCurrentScene().Transform.GetGlobalRect()
		var realRect common.RectBoundary
		var deltaX, deltaY = float32(event.Get("deltaX").Float()), float32(event.Get("deltaY").Float())
		//engine.LogDebug("deltaX, deltaY: %f %f", deltaX, deltaY)
		totalScrollX += deltaX
		if totalScrollX > realRect.X.Max{
			totalScrollX = realRect.X.Max
		}
		if totalScrollX < realRect.X.Min{
			totalScrollX = realRect.X.Min
		}

		totalScrollY += deltaY
		if totalScrollY > realRect.Y.Max{
			totalScrollY = realRect.Y.Max
		}
		if totalScrollY < realRect.Y.Min{
			totalScrollY = realRect.Y.Min
		}

		s.SceneObject.ForEachObject(func(object *engine.SceneObject) {
			size := object.Transform.GetGlobalRect()

			if size.X.Max - totalScrollX > realRect.X.Max {
				realRect.X.Max = size.X.Max - totalScrollX
			}
			if size.X.Min - totalScrollX < realRect.X.Min {
				realRect.X.Min = size.X.Min - totalScrollX
			}
			if size.Y.Max - totalScrollY > realRect.Y.Max {
				realRect.Y.Max = size.Y.Max - totalScrollY
			}
			if size.Y.Min - totalScrollY < realRect.Y.Min {
				realRect.Y.Min = size.Y.Min - totalScrollY
			}
		})

		if (viewRect.X.Max + deltaX) > realRect.X.Max {
			viewRect.X.Max = realRect.X.Max
			viewRect.X.Min = realRect.X.Max - viewRect.X.GetLength()
		} else {
			viewRect.X.Max += deltaX
			viewRect.X.Min += deltaX
		}

		if (viewRect.X.Min + deltaX) < realRect.X.Min {
			viewRect.X.Min = realRect.X.Min
			viewRect.X.Max = realRect.X.Min + viewRect.X.GetLength()
		} else {
			viewRect.X.Max += deltaX
			viewRect.X.Min += deltaX
		}

		if (viewRect.Y.Max + deltaY) > realRect.Y.Max {
			viewRect.Y.Max = realRect.Y.Max
			viewRect.Y.Min = realRect.Y.Max - viewRect.Y.GetLength()
		} else {
			viewRect.Y.Max += deltaY
			viewRect.Y.Min += deltaY
		}

		if (viewRect.Y.Min + deltaY) > realRect.Y.Min {
			viewRect.Y.Min = realRect.Y.Min
			viewRect.Y.Max = realRect.Y.Min + viewRect.Y.GetLength()
		} else{
			viewRect.Y.Max += deltaY
			viewRect.Y.Min += deltaY
		}
		engine.LogDebug("stopped at viewRect, realRect %+v, %+v", viewRect, realRect)

		s.SceneObject.Transform.Position.X, s.SceneObject.Transform.Position.Y = viewRect.X.Min, viewRect.Y.Min

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

func (s *Scrolling) OnStart(){
	s.ComponentImpl.OnStart()
}

func (s *Scrolling) OnUpdate(ctx engine.UpdateContext) {
}

func (s *Scrolling) GetName() string {
	return engine.NameOfComponent(s)
}

