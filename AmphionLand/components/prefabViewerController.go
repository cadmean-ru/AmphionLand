package components

import (
	"AmphionLand/generated/res"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type PrefabViewerController struct {
	engine.ComponentImpl
	PrefabPath string `state`
	Text string
	flag bool
}

func (s *PrefabViewerController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	if s.flag {
		s.SetTextView(s.Text)
	}

	s.SceneObject.GetChildByName("Text").GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView).SetHTextAlign(a.TextAlignCenter)
	s.SceneObject.GetChildByName("Text").GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView).SetVTextAlign(a.TextAlignCenter)
	s.SceneObject.GetChildByName("Text").GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView).SetFontSize(69)
}

func (s *PrefabViewerController) GetName() string {
	return engine.NameOfComponent(s)
}

func (s *PrefabViewerController) SetTextView(text string) {
	s.Text = text
	if s.SceneObject == nil {
		s.flag = true
		return
	}
	s.SceneObject.GetChildByName("Text").GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView).SetText(text)
}

func OnClick(event engine.AmphionEvent) bool {
	// path := event.Sender.(*engine.SceneObject).GetComponentByName("AmphionLand/components.PrefabViewerController").(*PrefabViewerController).PrefabPath
	leftScene := engine.GetInstance().GetCurrentScene().GetChildByName("left_scene")
	prefab, err := engine.LoadPrefab(res.Prefabs_button) //engine.GetInstance().GetResourceManager().IdOf(path))
	if err != nil {
		engine.LogDebug(err.Error())
		return false
	}
	leftScene.AddChild(prefab)
	return true
}