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

		pref.ForEachObject(func(object *engine.SceneObject) {
			for _, c := range object.GetComponents(true) {
				if _, isView := c.(engine.ViewComponent); !isView {
					engine.LogDebug("Removing component %v", c)
					pref.RemoveComponent(c)
				}
			}
		}, true)

		pref.AddComponent(&Yeeter{
			prefId: prefId,
		})

		for _, c := range pref.GetComponents(true) {
			engine.LogDebug(c.GetName())
		}

		engine.GetCurrentScene().AddChild(pref)
		engine.LogDebug("Here")
	}).Err(func(err error) {
		engine.LogDebug(err.Error())
	}).Build())

	return true
}