package components

import (
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
	currentPrefab := engine.NewSceneObject("temp")
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
		for _, box := range leftScene.GetChildren() {
			if len(box.GetChildren()) == 0 {
				pref :=res.(*engine.SceneObject)
				currentPrefab = pref
				box.AddChild(pref)
				pref.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,a.MatchParent)
				break
			}
		}
	}).Err(func(err error) {
		engine.LogDebug(err.Error())
	}).Build())

	hierarchy := engine.GetCurrentScene().GetChildByName("right_thing").GetChildByName("Hierarchy")
	hierarchy.ForEachChild(func(child *engine.SceneObject) {
		hierarchy.RemoveChild(child)
	})

	transMap := make(map[int]a.Vector3)
	transMap[0] = currentPrefab.Transform.Position
	transMap[1] = currentPrefab.Transform.Rotation
	transMap[2] = currentPrefab.Transform.Size

	for i := 0; i < 3; i++ {
		engine.LogDebug("loop")
		box := engine.NewSceneObject("emptyBox" + string(rune(i)))
		box.Transform.Size.Y = 90
		grid := builtin.NewGridLayout()
		grid.Cols = 2
		grid.Rows = 3
		box.AddComponent(grid)

		nameMap := make(map[int]string)
		nameMap[0] = "X"
		nameMap[1] = "Y"
		nameMap[2] = "Z"

		valueMap := make(map[int] float32)
		valueMap[0] = transMap[i].X
		valueMap[1] = transMap[i].Y
		valueMap[2] = transMap[i].Z

		for j := 0; j < 3; j++{
			name := engine.NewSceneObject("name" + string(rune(i)) + string(rune(j)))
			nameText := builtin.NewTextView(nameMap[j])
			nameText.SetVTextAlign(a.TextAlignCenter)
			nameText.SetHTextAlign(a.TextAlignCenter)

			name.AddComponent(nameText)
			name.Transform.Size.Y = 30
			name.AddComponent(builtin.NewShapeView(builtin.ShapeRectangle))

			value := engine.NewSceneObject("value" + string(rune(i)) + string(rune(j)))
			valueText := builtin.NewTextView(string(int32(valueMap[j])))
			valueText.SetVTextAlign(a.TextAlignCenter)
			valueText.SetHTextAlign(a.TextAlignCenter)

			value.AddComponent(valueText)
			value.Transform.Size.Y = 30
			value.AddComponent(builtin.NewShapeView(builtin.ShapeRectangle))

			box.AddChild(name)
			box.AddChild(value)
		}

		hierarchy.AddChild(box)
	}

	//engine.RunTask(engine.NewTaskBuilder().Run(func() (interface{}, error) {
	//	engine.LogDebug(path)
	//	id := engine.GetResourceManager().IdOf(path)
	//	engine.LogDebug("Id %d", id)
	//	return engine.GetResourceManager().ReadFile(id)
	//}).Then(func(res interface{}) {
	//	engine.LogDebug("%+v", string(res.([]byte)))
	//}).Err(func(err error) {
	//	engine.LogDebug(err.Error())
	//}).Build())

	return true
}