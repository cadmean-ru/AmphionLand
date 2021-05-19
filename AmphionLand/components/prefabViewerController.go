package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"strconv"
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

//func Yeet(event engine.AmphionEvent) bool {
//	senderObj := event.Sender.(*engine.SceneObject)
//	path := senderObj.GetComponentByName("AmphionLand/components.PrefabViewerController").(*PrefabViewerController).PrefabPath
//	engine.LogDebug("Here 2")
//
//	engine.RunTask(engine.NewTaskBuilder().Run(func() (interface{}, error) {
//		return engine.LoadPrefab(engine.GetResourceManager().IdOf(path))
//	}).Then(func(res interface{}) {
//		pref :=res.(*engine.SceneObject)
//		pref.AddComponent(&Yeeter{})
//		engine.GetCurrentScene().AddChild(pref)
//		engine.LogDebug("Here")
//	}).Err(func(err error) {
//		engine.LogDebug(err.Error())
//	}).Build())
//
//	return true
//}

func OnClick(event engine.AmphionEvent) bool {
	senderObj := event.Sender.(*engine.SceneObject)
	path := senderObj.GetComponentByName("AmphionLand/components.PrefabViewerController").(*PrefabViewerController).PrefabPath
	currentPrefab := engine.NewSceneObject("temp")

	prefId := engine.GetResourceManager().IdOf(path)

	engine.RunTask(engine.NewTaskBuilder().Run(func() (interface{}, error) {
		return engine.LoadPrefab(prefId)
	}).Then(func(res interface{}) {
		//leftScene := engine.GetCurrentScene().GetChildByName("left_scene")
		//for _, box := range leftScene.GetChildren() {
		//	if len(box.GetChildren()) == 0 {
		//		pref :=res.(*engine.SceneObject)
		//		currentPrefab = pref
		//		box.AddChild(pref)
		//		pref.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,a.MatchParent)
		//		break
		//	}
		//}

		pref := res.(*engine.SceneObject)

		pref.ForEachObject(func(object *engine.SceneObject) {
			for _, c := range object.GetComponents(true) {
				_, isView := c.(engine.ViewComponent)

				if !isView {
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

	hierarchy := engine.FindObjectByName("Hierarchy")
	hierarchy.ForEachChild(func(child *engine.SceneObject) {
		hierarchy.RemoveChild(child)
	})

	transMap := [3]a.Vector3{currentPrefab.Transform.Position, currentPrefab.Transform.Rotation, currentPrefab.Transform.Size}

	transform := [3]string{"Position", "Rotation", "Size"}

	for i := 0; i < 3; i++ {
		nameBox := engine.NewSceneObject("nameBox" + string(rune(i)))
		nameBox.Transform.Size.Y = 30
		nameBoxText := builtin.NewTextView(transform[i])
		nameBoxText.SetVTextAlign(a.TextAlignCenter)
		nameBoxText.SetHTextAlign(a.TextAlignCenter)
		nameBox.AddComponent(nameBoxText)
		hierarchy.AddChild(nameBox)

		box := engine.NewSceneObject("emptyBox" + string(rune(i)))
		box.Transform.Size.Y = 90
		grid := builtin.NewGridLayout()
		grid.Cols = 2
		grid.Rows = 3
		box.AddComponent(grid)

		nameMap := [3]string{"X", "Y", "Z"}
		valueMap := [3]float32{transMap[i].X, transMap[i].Y, transMap[i].Z}

		for j := 0; j < 3; j++{
			name := engine.NewSceneObject("name" + string(rune(i)) + string(rune(j)))
			nameText := builtin.NewTextView(nameMap[j])
			nameText.SetVTextAlign(a.TextAlignCenter)
			nameText.SetHTextAlign(a.TextAlignCenter)

			name.AddComponent(nameText)
			name.Transform.Size.Y = 30

			value := engine.NewSceneObject("value" + string(rune(i)) + string(rune(j)))
			valueText := builtin.NewTextView(strconv.FormatFloat(float64(valueMap[j]), 'f',3,32))
			valueText.SetVTextAlign(a.TextAlignCenter)
			valueText.SetHTextAlign(a.TextAlignCenter)

			value.AddComponent(valueText)
			value.Transform.Size.Y = 30
			//value.AddComponent(builtin.NewShapeView(builtin.ShapeRectangle))

			box.AddChild(name)
			box.AddChild(value)
		}

		hierarchy.AddChild(box)
	}

	return true
}