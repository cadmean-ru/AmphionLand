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
}

func (s *PrefabViewerController) OnStart() {
	textObj := s.SceneObject.GetChildByName("Text")
	textObj.GetComponentByName("TextView", true).(*builtin.TextView).SetText(s.Text)
}

func (s *PrefabViewerController) GetName() string {
	return engine.NameOfComponent(s)
}

func OnClick(event engine.AmphionEvent) bool {
	senderObj := event.Sender.(*engine.SceneObject)
	path := senderObj.GetComponentByName("AmphionLand/components.PrefabViewerController").(*PrefabViewerController).PrefabPath

	//prefab, err := engine.LoadPrefab(res.Prefabs_button) //engine.GetInstance().GetResourceManager().IdOf(path))
	//if err != nil {
	//	engine.LogError(err.Error())
	//	return false
	//}
	//leftScene.AddChild(prefab)
	engine.RunTask(engine.NewTaskBuilder().Run(func() (interface{}, error) {
		return engine.LoadPrefab(engine.GetResourceManager().IdOf(path))
	}).Then(func(res interface{}) {
		leftScene := engine.GetCurrentScene().GetChildByName("left_scene")
		leftScene.AddChild(res.(*engine.SceneObject))
	}).Err(func(err error) {
		engine.LogDebug(err.Error())
	}).Build())

	return true
}