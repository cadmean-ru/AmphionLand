package components

import (
	"github.com/cadmean-ru/amphion/engine"
)

type EmptyBox struct {
	engine.ComponentImpl
	xMin int `state`
	yMin int `state`
	xMax int `state`
	yMax int `state`
	isAvailable bool `state`
}

func (s *EmptyBox) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
}

func (s *EmptyBox) GetName() string {
	return engine.NameOfComponent(s)
}
