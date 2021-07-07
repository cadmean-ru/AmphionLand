package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"math"
)

type EditorGrid struct {
    engine.ComponentImpl
    grid *builtin.GridLayout
    editor *EditorController
}

func (s *EditorGrid) OnStart() {
	s.grid = s.SceneObject.GetComponentByName("GridLayout").(*builtin.GridLayout)
	s.editor = engine.FindComponentByName("EditorController").(*EditorController)

	for i := 0; i < s.grid.GetRowsCount() * s.grid.GetColumnsCount(); i++ {
		s.SceneObject.AddChild(s.makeBox(i, false))
	}
}

func (s *EditorGrid) MakeClickable() {
	s.SceneObject.ForEachChild(func(object *engine.SceneObject) {
		object.AddComponent(builtin.NewRectBoundary())
	})
}

func (s *EditorGrid) AdjustSize(x, y int) {
	currentSize := s.grid.GetColumnsCount() * s.grid.GetRowsCount()
	newSize := x * y
	diff := int(math.Abs(float64(newSize - currentSize)))

	children := s.SceneObject.GetChildren()
	for i := 0; i < diff; i++ {
		if newSize > currentSize {
			s.SceneObject.AddChild(s.makeBox(currentSize + i, true))
		} else {
			s.SceneObject.RemoveChild(children[len(children) - 1 - i])
		}
	}
}

func (s *EditorGrid) makeBox(num int, addBoundary bool) *engine.SceneObject {
	box := engine.NewSceneObject(fmt.Sprintf("Box %d", num))
	rectangle := builtin.NewShapeView(builtin.ShapeRectangle)
	rectangle.FillColor = a.TransparentColor()
	rectangle.StrokeColor = a.BlackColor()
	if s.editor.showBoxed {
		rectangle.StrokeWeight = 1
	} else {
		rectangle.StrokeWeight = 0
	}
	box.Transform.Size.Y = 100
	box.AddComponent(rectangle)
	if addBoundary {
		box.AddComponent(builtin.NewRectBoundary())
	}
	box.AddComponent(&ClickAndInspeceet{})
	return box
}

func (s *EditorGrid) GetName() string {
    return engine.NameOfComponent(s)
}