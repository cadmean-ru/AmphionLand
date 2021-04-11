package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
)

type WodgetController struct {
	engine.ComponentImpl
	wodget *engine.SceneObject
	moving bool
	mousePos a.IntVector2
}

func (s *WodgetController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
	engine.LogDebug("OnInit")
}

func (s *WodgetController) GetName() string {
	return engine.NameOfComponent(s)
}
