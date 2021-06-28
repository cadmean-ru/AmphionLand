package components

import (
	"github.com/cadmean-ru/amphion/common"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"math"
)

type ScrollDirection byte

const (
	ScrollUp ScrollDirection = iota
	ScrollDown
	ScrollNone
)

type NewNewScrollManager struct {
	engine.ComponentImpl
	dScrollY         float32
	scrollDirectionY ScrollDirection
}

func (s *NewNewScrollManager) OnStart() {
	engine.BindEventHandler(engine.EventMouseScroll, s.handleScroll)
}

func (s *NewNewScrollManager) OnStop() {
	engine.UnbindEventHandler(engine.EventMouseScroll, s.handleScroll)
}

//func (s *NewNewScrollManager) LayoutChildren() {
//
//}

func (s *NewNewScrollManager) handleScroll(event engine.AmphionEvent) bool {
	scrollAmount := event.Data.(a.Vector2)

	if scrollAmount.Y < 0 {
		s.scrollDirectionY = ScrollUp
	} else if scrollAmount.Y > 0 {
		s.scrollDirectionY = ScrollDown
	} else {
		s.scrollDirectionY = ScrollNone
	}

	s.dScrollY = scrollAmount.Y

	var realArea = s.measureChildren()
	var sceneRect = engine.GetCurrentScene().Transform.GetRect()

	//engine.LogDebug("%+v", realArea)

	var theScrolly float32

	if s.scrollDirectionY == ScrollUp {
		if realArea.GetMin().Y < 0 {
			mouseScroll := float64(s.dScrollY)
			areaOffset := float64(realArea.GetMin().Y)
			//engine.LogDebug("%f %f", mouseScroll, areaOffset)
			theScrolly = float32(math.Max(mouseScroll, areaOffset))
		}
	} else if s.scrollDirectionY == ScrollDown {
		if realArea.GetMax().Y > sceneRect.Y.GetLength() {
			mouseScroll := float64(s.dScrollY)
			areaOffset := float64(realArea.GetMax().Y - sceneRect.Y.GetLength())
			//engine.LogDebug("%f %f %f", mouseScroll, areaOffset, sceneRect.Y.GetLength())
			theScrolly = float32(math.Min(mouseScroll, areaOffset))
		}
	}

	//engine.LogDebug("Scrolling %d %f", s.scrollDirectionY, theScrolly)

	if s.scrollDirectionY != ScrollNone {
		s.SceneObject.Transform.Position = s.SceneObject.Transform.Position.Sub(a.NewVector3(0, theScrolly, 0))

		s.SceneObject.ForEachObject(func(object *engine.SceneObject) {
			object.Redraw()
		})
		s.SceneObject.Redraw()
	}

	s.scrollDirectionY = ScrollNone

	return true
}

func (s *NewNewScrollManager) measureChildren() *common.RectBoundary {
	realArea := common.NewRectBoundary(0, 0, 0, 0, -999, 999)

	s.SceneObject.ForEachObject(func(object *engine.SceneObject) {
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

	return realArea
}

func (s *NewNewScrollManager) GetName() string {
	return engine.NameOfComponent(s)
}

func NewNewNewScrollManager() *NewNewScrollManager {
	return &NewNewScrollManager{}
}
