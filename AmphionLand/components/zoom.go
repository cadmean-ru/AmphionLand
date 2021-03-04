package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type Zooming struct {
	engine.ComponentImpl
	ZoomNumber int `state`
	EventOn int `state`
	EventOff int `state`
}

func (s *Zooming) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
	s.SceneObject.AddComponent(builtin.NewEventListener(s.EventOn,
		func(mouseHoverEvent engine.AmphionEvent) bool{
			for _, component:= range s.SceneObject.GetComponentsByName(".*TextView"){
				if view, ok := component.(*builtin.TextView); ok {
					view.FontSize *= byte(s.ZoomNumber)
					view.ForceRedraw()
					engine.RequestRendering()
				}
			}
			s.SceneObject.SetSize(s.SceneObject.Transform.Size.Multiply(a.Vector3{X: float32(s.ZoomNumber),
				Y: float32(s.ZoomNumber), Z: float32(s.ZoomNumber)}))
			return true
		}))
	s.SceneObject.AddComponent(builtin.NewEventListener(s.EventOff,
		func(mouseHoverEvent engine.AmphionEvent) bool{
			for _, component:= range s.SceneObject.GetComponentsByName(".*TextView"){
				if view, ok := component.(*builtin.TextView); ok {
					view.FontSize /= byte(s.ZoomNumber)
					view.ForceRedraw()
					engine.RequestRendering()
				}
			}
			s.SceneObject.SetSize(s.SceneObject.Transform.Size.Multiply(a.Vector3{X: 1/float32(s.ZoomNumber),
				Y: 1/float32(s.ZoomNumber), Z: 1/float32(s.ZoomNumber)}))
			return true
		}))
}

func (s *Zooming) GetName() string {
	return engine.NameOfComponent(s)
}
