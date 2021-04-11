package components

import (
	"github.com/cadmean-ru/amphion/engine"
)

type Wodget struct {
	engine.ComponentImpl
	heightNumber int `state`
	widthNumber int `state`
}

func (s *Wodget) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
	s.heightNumber = 1
	s.widthNumber = 2
}

func (s *Wodget) GetName() string {
	return engine.NameOfComponent(s)
}
