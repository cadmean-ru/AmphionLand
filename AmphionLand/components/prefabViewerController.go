package components

import (
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
