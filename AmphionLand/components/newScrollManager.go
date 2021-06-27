package components

import (
	"github.com/cadmean-ru/amphion/common"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"math"
	"strings"
)

type NewScrolling struct {
	engine.ComponentImpl
	pipiRect *engine.SceneObject
}

func (s *NewScrolling) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	s.pipiRect = engine.NewSceneObject("pipiRect")
	pipiRectView := builtin.NewShapeView(builtin.ShapeRectangle)
	pipiRectView.FillColor = a.TransparentColor()
	pipiRectView.StrokeWeight = 10
	pipiRectView.StrokeColor = a.RedColor()
	s.pipiRect.AddComponent(pipiRectView)
	engine.GetCurrentScene().AddChild(s.pipiRect)
}

func (s *NewScrolling) OnStart(){
	s.ComponentImpl.OnStart()

	var offset a.Vector3
	engine.BindEventHandler(engine.EventMouseScroll, func(event engine.AmphionEvent) bool {
		o := event.Data.(a.Vector2)
		dOffset := a.NewVector3(o.X, o.Y, 0)

		//engine.LogDebug("Scroll: %f %f", dOffset.X, dOffset.Y)

		scene := engine.GetCurrentScene()
		ss := scene.Transform.GetSize()
		visibleArea := common.NewRectBoundary(-offset.X, -offset.X + ss.X, -offset.Y, -offset.Y + ss.Y, -999, 999)

		s.pipiRect.Transform.Position = visibleArea.GetMin()
		s.pipiRect.Transform.Position.Z = 999
		s.pipiRect.Transform.Size = visibleArea.GetSize()

		//engine.LogDebug("va %f %f", visibleArea.Y.Min, visibleArea.Y.Max)

		//realArea := common.NewRectBoundary(0, 0, 0, 0, -999, 999)
		//scene.ForEachObject(func(object *engine.SceneObject) {
		//	rect := object.Transform.GetGlobalRect()
		//	if rect.X.Min < realArea.X.Min {
		//		realArea.X.Min = rect.X.Min
		//	}
		//	if rect.X.Max > realArea.X.Max {
		//		realArea.X.Max = rect.X.Max
		//	}
		//	if rect.Y.Min < realArea.Y.Min {
		//		realArea.Y.Min = rect.Y.Min
		//	}
		//	if rect.Y.Max > realArea.Y.Max {
		//		realArea.Y.Max = rect.Y.Max
		//	}
		//})

		var scrollingDown, scrollingUp = dOffset.Y < 0, dOffset.Y > 0
		var canScrollDown, canScrollUp bool
		var minOutY float32 = -1


		scene.ForEachObject(func(object *engine.SceneObject) {
			if strings.Contains(object.GetName(), "Box"){
				engine.LogDebug("va %f %f", visibleArea.Y.Min, visibleArea.Y.Max)
				//engine.LogDebug("%f %f \n" + object.GetName(), object.Transform.GetGlobalRect().Y.Min, object.Transform.GetGlobalRect().Y.Max)
			}

			rect := object.Transform.GetGlobalRect()
			if !visibleArea.IsRectInside(rect) {
				if strings.Contains(object.GetName(), "Box"){
					engine.LogDebug("%f %f \n" + object.GetName(), object.Transform.GetGlobalRect().Y.Min, object.Transform.GetGlobalRect().Y.Max)
					//engine.LogDebug(object.GetName())
				}

				if scrollingDown && !canScrollDown { // down
					canScrollDown = rect.Y.Min > visibleArea.Y.Max || rect.Y.Max > visibleArea.Y.Max
									//519			 73					  619		   73
					if canScrollDown {
						//if strings.Contains(object.GetName(), "Box"){
						//	engine.LogDebug("%f %f \n" + object.GetName() + " csd", object.Transform.GetGlobalRect().Y.Min, object.Transform.GetGlobalRect().Y.Max)
							engine.LogDebug(object.GetName() + " csd")
						//}
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
		engine.LogDebug("\n")

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

		engine.LogDebug("dOffset: %+v; old offset: %+v", dOffset, offset)

		offset = offset.Add(dOffset)
		scene.Transform.Position = offset
		engine.LogDebug("new offset: %+v", offset)

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

