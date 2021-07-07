package components

import (
	"AmphionLand/res"
	"github.com/cadmean-ru/amphion/engine"
)

type NewSceneController struct {
	engine.ComponentImpl

}

func (s *NewSceneController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	inputObj, _ := engine.LoadPrefab(res.Prefabs_inputBox)
	s.SceneObject.AddChild(inputObj)

}

func (s *NewSceneController) OnStart() {
	engine.LogDebug("New OnStart")
}

func (s *NewSceneController) GetName() string {
	return engine.NameOfComponent(s)
}