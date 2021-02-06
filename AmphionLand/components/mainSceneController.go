//+build js

package components

import (
	"AmphionLand/res"
	"github.com/cadmean-ru/amphion/engine"
)

type MainSceneController struct {
	engine.ComponentImpl
}

func (s *MainSceneController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	inputBoxPrefab, ex := engine.LoadPrefab(res.Prefabs_inputBox)
	if ex == nil {
		s.SceneObject.AddChild(inputBoxPrefab)
	}
}

func (s *MainSceneController) GetName() string {
	return engine.NameOfComponent(s)
}