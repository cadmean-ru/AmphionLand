package test

import "github.com/cadmean-ru/amphion/engine"

type Test35 struct {
    engine.ComponentImpl
}

func (s *Test35) OnStart() {

}

func (s *Test35) GetName() string {
    return engine.NameOfComponent(s)
}

func NewTest35() *Test35 {
    return &Test35{}
}
