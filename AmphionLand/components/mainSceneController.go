//+build js

package components

import (
	"AmphionLand/generated/res"
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

	searchBarPrefab, ex := engine.LoadPrefab(res.Prefabs_searchBar)
	if ex == nil {
		s.SceneObject.AddChild(searchBarPrefab)
	} else {
		engine.LogDebug(ex.Error())
	}
}

func (s *MainSceneController) OnStart() {
	engine.Navigate("login", nil)
}

func (s *MainSceneController) GetName() string {
	return engine.NameOfComponent(s)
}