package components

import (
	"github.com/cadmean-ru/amphion/common"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"math"
)

type NewScrolling struct {
	engine.ComponentImpl
}

func (s *NewScrolling) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

}

func (s *NewScrolling) OnStart(){
	s.ComponentImpl.OnStart()

	var offset a.Vector3
	engine.BindEventHandler(engine.EventMouseScroll, func(event engine.AmphionEvent) bool {
		o := event.Data.(a.Vector2)
		dOffset := a.NewVector3(o.X, o.Y, 0)

		engine.LogDebug("Scroll: %f %f", dOffset.X, dOffset.Y)

		scene := engine.GetCurrentScene()
		ss := scene.Transform.GetSize()
		visibleArea := common.NewRectBoundary(-offset.X, -offset.X + ss.X, -offset.Y, -offset.Y + ss.Y, -999, 999)
		realArea := common.NewRectBoundary(0, 0, 0, 0, -999, 999)
		scene.ForEachObject(func(object *engine.SceneObject) {
			rect := object.Transform.GetGlobalRect()
			if rect.X.Min < realArea.X.Min {
				realArea.X.Min = rect.X.Min
			}
			if rect.X.Max > realArea.X.Max {
				realArea.X.Max = rect.X.Max
			}
			if rect.Y.Min < realArea.Y.Min {
				realArea.Y.Min = rect.Y.Min
			}
			if rect.Y.Max > realArea.Y.Max {
				realArea.Y.Max = rect.Y.Max
			}
		})

		var scrollingDown, scrollingUp = dOffset.Y < 0, dOffset.Y > 0
		var canScrollDown, canScrollUp bool
		var minOutY float32 = -1

		scene.ForEachObject(func(object *engine.SceneObject) {
			rect := object.Transform.GetGlobalRect()
			if !visibleArea.IsRectInside(rect) {
				//if dOffset.Y > 0 { // if scrolling up
				//	if rect.Y.Min < visibleArea.Y.Min {
				//		// can scroll
				//		//finalDY = common.ClampFloat32(dOffset.Y, -visibleArea.Y.Min + rect.Y.Min, 0)
				//		dOutY := visibleArea.Y.Min - rect.Y.Min
				//		if minOutY1 == -1 || dOutY < minOutY1 {
				//			minOutY1 = dOutY
				//		}
				//	}
				//} else if dOffset.Y < 0 { // scrolling down
				//	if rect.Y.Max > visibleArea.Y.Max {
				//		// can scroll
				//		//finalDY = common.ClampFloat32(dOffset.Y, -visibleArea.Y.Min + rect.Y.Min, 0)
				//		dOutY := rect.Y.Max - visibleArea.Y.Max
				//		if minOutY2 == -1 || dOutY < minOutY2 {
				//			minOutY2 = dOutY
				//		}
				//	}
				//}

				if scrollingDown && !canScrollDown { // down
					canScrollDown = rect.Y.Min > visibleArea.Y.Max || rect.Y.Max > visibleArea.Y.Max
					if canScrollDown {
						m := float32(math.Min(math.Abs(float64(rect.Y.Min-visibleArea.Y.Max)), math.Abs(float64(rect.Y.Max-visibleArea.Y.Max))))
						if minOutY == -1 || m < minOutY {
							minOutY = m
						}
					}
				} else if scrollingUp && !canScrollUp { // up
					canScrollUp = rect.Y.Min < visibleArea.Y.Min || rect.Y.Max < visibleArea.Y.Min
					if canScrollUp {
						m := float32(math.Min(math.Abs(float64(visibleArea.Y.Min - rect.Y.Min)), math.Abs(float64(visibleArea.Y.Min - rect.Y.Max))))
						if minOutY == -1 || m < minOutY {
							minOutY = m
						}
					}
				}
			}
		})

		if scrollingDown {
			if !canScrollDown{
				dOffset.Y = 0
			} else {
				dOffset.Y = float32(math.Min(float64(dOffset.Y), float64(-minOutY)))
			}
		}
		if scrollingUp {
			if !canScrollUp{
				dOffset.Y = 0
			} else {
				dOffset.Y = float32(math.Min(float64(dOffset.Y), float64(minOutY)))
			}
		}

		scene.Transform.Position = scene.Transform.Position.Add(dOffset)
		offset = offset.Add(dOffset)
		engine.ForceAllViewsRedraw()
		engine.RequestRendering()
		return true
	})
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

