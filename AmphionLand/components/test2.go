package components

import "github.com/cadmean-ru/amphion/engine"

type Test2 struct {
	engine.ViewImpl
	Text float64 `state:"text69"`
}

func (s *Test2) OnInit(ctx engine.InitContext) {
	s.ViewImpl.OnInit(ctx)
}

func (s *Test2) OnStart() {

}

func (s *Test2) GetName() string {
	return engine.NameOfComponent(s)
}
