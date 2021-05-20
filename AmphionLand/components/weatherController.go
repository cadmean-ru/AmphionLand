package components

import "github.com/cadmean-ru/amphion/engine"

type WeatherController struct {
	engine.ComponentImpl
}

func (s *WeatherController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
}

func (s *WeatherController) GetName() string {
	return engine.NameOfComponent(s)
}