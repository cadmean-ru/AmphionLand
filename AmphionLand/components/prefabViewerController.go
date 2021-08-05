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
	textObj.GetComponentByName("TextView").(*builtin.TextView).SetText(s.Text)
}

func (s *PrefabViewerController) GetName() string {
	return engine.NameOfComponent(s)
}

func OnClick(event engine.AmphionEvent) bool {
	senderObj := event.Sender.(*engine.SceneObject)
	path := senderObj.GetComponentByName("AmphionLand/components.PrefabViewerController").(*PrefabViewerController).PrefabPath
	//currentPrefab := engine.NewSceneObject("temp")
	pref := engine.NewSceneObject("temp")

	prefId := engine.GetResourceManager().IdOf(path)

	engine.RunTask(engine.NewTaskBuilder().Run(func() (interface{}, error) {
		return engine.LoadPrefab(prefId)
	}).Then(func(res interface{}) {
		pref = res.(*engine.SceneObject)
		engine.LogDebug("name " + pref.GetName())

		preparePrefabForEditor(pref)

		pref.AddComponent(&Yeeter{
			prefId: prefId,
		})

		engine.GetCurrentScene().AddChild(pref)
		engine.LogDebug("Here")
	}).Err(func(err error) {
		engine.LogDebug(err.Error())
	}).Build())

	return true
}

func preparePrefabForEditor(pref *engine.SceneObject) {
	engine.LogDebug("preparing %s", pref.GetName())

	pref.ForEachObject(func(object *engine.SceneObject) {
		for _, c := range object.GetComponents(true) {
			if componentIsNotNecessary(c) {
				engine.LogDebug("Removing component %s", engine.NameOfComponent(c))
				object.RemoveComponent(c)
			}
		}
	}, true)
}

func componentIsNotNecessary(c engine.Component) bool {
	_, isView := c.(engine.ViewComponent)
	_, isLayout := c.(engine.Layout)

	return !engine.ComponentNameMatches(engine.NameOfComponent(c), "EditorGrid") && !isView && !isLayout
}