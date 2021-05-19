package components

import (
	"AmphionLand/generated/res"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"strconv"
)

type ClickAndInspeceet struct {
	engine.ComponentImpl
}

func (s *ClickAndInspeceet) OnStart() {
	engine.LogDebug("inspeceet start")
	engine.BindEventHandler(engine.EventMouseDown, s.handleClick)
}

func (s *ClickAndInspeceet) OnStop() {
	engine.UnbindEventHandler(engine.EventMouseDown, s.handleClick)
}

func (s *ClickAndInspeceet) GetName() string {
	return engine.NameOfComponent(s)
}

func (s *ClickAndInspeceet) handleClick(event engine.AmphionEvent) bool {
	s.showInspectorForObject(event.Data.(engine.MouseEventData).SceneObject)
	return false
}

func (s *ClickAndInspeceet) showInspectorForObject(object *engine.SceneObject) {
	hierarchy := engine.FindObjectByName("Hierarchy")
	hierarchy.RemoveAllChildren()

	objectNameBox := engine.NewSceneObject("objectNameBox")
	objectNameBox.Transform.Size.Y = 30
	objectNameBoxText := builtin.NewTextView(object.GetName())
	objectNameBoxText.SetVTextAlign(a.TextAlignCenter)
	objectNameBoxText.SetHTextAlign(a.TextAlignCenter)
	objectNameBox.AddComponent(objectNameBoxText)
	hierarchy.AddChild(objectNameBox)

	transMap := map[string]a.Vector3 {
		"Position": object.Transform.Position,
		"Rotation": object.Transform.Rotation,
		"Size":     object.Transform.Size,
	}

	i := 0
	for vecName, vec := range transMap {
		nameBox := engine.NewSceneObject("nameBox" + string(rune(i)))
		nameBox.Transform.Size.Y = 30
		nameBoxText := builtin.NewTextView(vecName)
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
		valueMap := [3]float32{vec.X, vec.Y, vec.Z}

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

		i++
	}

	if object.GetName()=="horizontalYeeter"{
		gridObject := engine.NewSceneObject("grid bruh")
		colsAmount := object.GetComponentByName("GridLayout").(*builtin.GridLayout).Rows

		grid := builtin.NewGridLayout()
		grid.Cols = 2
		grid.Rows = 1

		gridObject.AddComponent(grid)

		inputObj, _ := engine.LoadPrefab(res.Prefabs_inputBox)
		buttonObj, _ := engine.LoadPrefab(res.Prefabs_button)

		inputText := inputObj.FindComponentByName("TextView").(*builtin.TextView)
		inputText.SetText(strconv.Itoa(colsAmount))

		buttonObj.FindComponentByName("TextView").(*builtin.TextView).SetText("ok")

		buttonObj.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
			inputObjText,_ := strconv.Atoi(inputText.GetText())
			object.FindComponentByName("GridLayout").(*builtin.GridLayout).Cols = inputObjText
			return true
		}))

		gridObject.AddChild(inputObj)
		gridObject.AddChild(buttonObj)

		hierarchy.AddChild(gridObject)
	}
}
