package components

import (
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type EditorGrid struct {
    engine.ComponentImpl
    grid *builtin.GridLayout
}

func (s *EditorGrid) OnStart() {
	s.grid = s.SceneObject.GetComponentByName("GridLayout").(*builtin.GridLayout)

	for i := 0; i < s.grid.Rows * s.grid.Cols; i++ {
		s.SceneObject.AddChild()
	}
}

func (s *EditorGrid) GetName() string {
    return engine.NameOfComponent(s)
}

func NewEditorGrid() *EditorGrid {
    return &EditorGrid{}
}
