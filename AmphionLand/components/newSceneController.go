package components

import (
	"AmphionLand/generated/res"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
)

type NewSceneController struct {
	engine.ComponentImpl

}

func (s *NewSceneController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	inputObj, _ := engine.LoadPrefab(res.Prefabs_inputBox)
	inputObj.Transform.Size.X = 1000
	inputObj.Transform.Size.Y = 1000
	inputObj.Transform.Position = a.NewVector3(-5,-5,0)
	s.SceneObject.AddChild(inputObj)

	engine.GetCurrentScene().AddComponent(NewNewNewScrollManager())

}

func (s *NewSceneController) OnStart() {
	engine.LogDebug("New OnStart")
}

func (s *NewSceneController) GetName() string {
	return engine.NameOfComponent(s)
}