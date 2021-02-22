//+build js

package components

import (
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)


type WidgetGrid struct {
	builtin.GridLayout
	Size int `state`
}

func (s *WidgetGrid) OnInit(ctx engine.InitContext) {
	s.GridLayout.OnInit(ctx)
	s.Size = 5
	s.Cols = s.Size
	s.Rows = s.Size
}

func (s *WidgetGrid) GetName() string {
	return engine.NameOfComponent(s)
}
