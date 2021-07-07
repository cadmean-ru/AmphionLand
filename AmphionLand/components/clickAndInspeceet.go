package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"strconv"
)

type ClickAndInspeceet struct {
	engine.ComponentImpl
	editor *EditorController
	hierarchy *engine.SceneObject
	shape *builtin.ShapeView
}

func (s *ClickAndInspeceet) OnStart() {
	engine.LogDebug("inspeceet start")
	engine.BindEventHandler(engine.EventMouseDown, s.handleClick)

	s.editor = engine.FindComponentByName("EditorController").(*EditorController)
	s.hierarchy = engine.FindObjectByName("Hierarchy")
	s.shape = s.SceneObject.GetComponentByName("ShapeView").(*builtin.ShapeView)
}

func (s *ClickAndInspeceet) OnStop() {
	engine.UnbindEventHandler(engine.EventMouseDown, s.handleClick)
}

func (s *ClickAndInspeceet) GetName() string {
	return engine.NameOfComponent(s)
}

func (s *ClickAndInspeceet) LayoutChildren() {
	children := s.SceneObject.GetChildren()
	if len(children) == 0 {
		return
	}

	first := children[0]
	first.Transform.Position = a.NewVector3(0, 0, 1)
	first.Transform.Pivot = a.ZeroVector()
	first.Transform.Size = a.NewVector3(a.MatchParent, a.MatchParent, a.MatchParent)

	for i := 1; i < len(children); i++ {
		c := children[i]
		c.Transform.Position = a.ZeroVector()
		c.Transform.Size = a.ZeroVector()
	}
}

func (s *ClickAndInspeceet) ToggleBox() {
	showBox := s.editor.showBoxed

	if showBox {
		s.shape.StrokeWeight = 1
	} else {
		s.shape.StrokeWeight = 0
	}

	s.shape.Redraw()
}

func (s *ClickAndInspeceet) handleClick(event engine.AmphionEvent) bool {
	if event.Data.(engine.MouseEventData).SceneObject != s.SceneObject {
		return true
	}

	engine.LogDebug("bruh 1")

	if s.editor.yeetingSceneObject == nil {
		if s.editor.remover {
			s.SceneObject.RemoveAllChildren()
			s.editor.hierarchy.RemoveAllChildren()
		} else if s.SceneObject.GetChildrenCount() == 1 {
			s.showInspector(s.SceneObject)
		} else {
			s.editor.hierarchy.RemoveAllChildren()
		}

		engine.RequestRendering()
		return false
	}

	engine.LogDebug("bruh 2")

	s.SceneObject.RemoveAllChildren()

	newObj := s.editor.yeetingSceneObject
	newObj.RemoveComponentByName("Yeeter")
	newObj.SetParent(s.SceneObject)
	s.editor.yeetingSceneObject = nil

	engine.LogDebug("here 1lekj")

	if newObj.GetName() == "Horizontal grid" {
		editorGrid := newObj.GetComponentByName("EditorGrid").(*EditorGrid)
		editorGrid.MakeClickable()
	}

	s.showInspector(s.SceneObject)

	engine.RequestRendering()
	return false
}

func (s *ClickAndInspeceet) showInspector(object *engine.SceneObject) {
	object = object.GetChildren()[0]
	engine.LogDebug(object.GetName())
	s.hierarchy.RemoveAllChildren()

	objectNameBox := engine.NewSceneObject("objectNameBox")
	objectNameBox.Transform.Size.Y = 30
	objectNameBoxText := builtin.NewTextView(object.GetName())
	objectNameBoxText.SetVTextAlign(a.TextAlignCenter)
	objectNameBoxText.SetHTextAlign(a.TextAlignCenter)
	objectNameBox.AddComponent(objectNameBoxText)
	s.hierarchy.AddChild(objectNameBox)

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
		s.hierarchy.AddChild(nameBox)

		box := engine.NewSceneObject(fmt.Sprintf("emptyBox %d", i))
		box.Transform.Size.Y = 90
		grid := builtin.NewGridLayout()
		grid.AddColumn(a.WrapContent)
		grid.AddColumn(a.WrapContent)
		grid.AddRow(a.WrapContent)
		grid.AddRow(a.WrapContent)
		grid.AddRow(a.WrapContent)

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
		s.hierarchy.AddChild(box)

		i++
	}

	components := object.GetComponents()
	cm := engine.GetInstance().GetComponentsManager()
	for _, comp := range components {

		componentsSomething := engine.NewSceneObject("King of components " + comp.GetName())

		publics := cm.GetComponentState(comp)
		componentsSomething.Transform.Size.Y = float32(30 * (len(publics) + 1))
		grid := builtin.NewGridLayout()
		grid.AddColumn(a.WrapContent)
		grid.AddColumn(a.WrapContent)

		componentsSomething.AddComponent(grid)

		nameBoxLabel := engine.NewSceneObject("nameBoxLabel")
		nameBoxLabel.Transform.Size.Y = 30
		nameBoxLabelText := builtin.NewTextView("Name: ")
		nameBoxLabelText.SetVTextAlign(a.TextAlignCenter)
		nameBoxLabelText.SetHTextAlign(a.TextAlignLeft)
		nameBoxLabel.AddComponent(nameBoxLabelText)
		componentsSomething.AddChild(nameBoxLabel)

		nameBox := engine.NewSceneObject("nameBox " + comp.GetName())
		nameBox.Transform.Size.Y = 30
		nameBoxText := builtin.NewTextView(comp.GetName())
		nameBoxText.SetVTextAlign(a.TextAlignCenter)
		nameBoxText.SetHTextAlign(a.TextAlignRight)
		nameBox.AddComponent(nameBoxText)
		componentsSomething.AddChild(nameBox)

		s.hierarchy.AddChild(componentsSomething)

		for name, public := range publics {
			stateLabel := engine.NewSceneObject("stateLabel " + name)
			stateLabel.Transform.Size.Y = 30
			stateLabelText := builtin.NewTextView(name)
			stateLabelText.SetHTextAlign(a.TextAlignLeft)
			stateLabelText.SetVTextAlign(a.TextAlignCenter)
			stateLabel.AddComponent(stateLabelText)
			componentsSomething.AddChild(stateLabel)

			stateInput := engine.NewSceneObject("stateInput" + name)
			stateInput.Transform.Size.Y = 30
			stateLabelInput := builtin.NewTextView(fmt.Sprintf("%+v", public))
			stateLabelInput.SetHTextAlign(a.TextAlignRight)
			stateLabelInput.SetVTextAlign(a.TextAlignCenter)
			stateInput.AddComponent(stateLabelInput)
			componentsSomething.AddChild(stateInput)

			switch public.(type) {
			case string:

			}

			engine.LogDebug("pub %s %v", name, public)
		}
	}

	//if object.GetName() == "Horizontal grid" {
	//	engine.LogDebug("clicked on hg")
	//	gridObject := engine.NewSceneObject("grid bruh")
	//
	//	grid := builtin.NewGridLayout()
	//	grid.Cols = 2
	//	grid.Rows = 1
	//
	//	gridObject.AddComponent(grid)
	//
	//	inputObj, _ := engine.LoadPrefab(res.Prefabs_inputBox)
	//	buttonObj, _ := engine.LoadPrefab(res.Prefabs_button)
	//
	//	actualGrid := object.GetComponentByName("GridLayout", true).(*builtin.GridLayout)
	//
	//	inputText := inputObj.FindComponentByName("TextView", true).(*builtin.TextView)
	//	inputText.SetText(strconv.Itoa(actualGrid.Cols))
	//
	//	buttonObj.FindComponentByName("TextView", true).(*builtin.TextView).SetText("ok")
	//
	//	buttonObj.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
	//		newCols,_ := strconv.Atoi(inputText.GetText())
	//		engine.LogDebug("Old cols: %d. New cols: %d", grid.Cols, newCols)
	//		actualGrid.Cols = newCols
	//		return true
	//	}))
	//
	//	gridObject.AddChild(inputObj)
	//	gridObject.AddChild(buttonObj)
	//
	//	s.hierarchy.AddChild(gridObject)
	//}
	//
	//if object.GetName() == "Text Label" {
	//	engine.LogDebug("clicked on textlabel")
	//	gridObject := engine.NewSceneObject("label bruh")
	//
	//	grid := builtin.NewGridLayout()
	//	grid.Cols = 2
	//	grid.Rows = 1
	//
	//	gridObject.AddComponent(grid)
	//
	//	inputObj, _ := engine.LoadPrefab(res.Prefabs_inputBox)
	//	buttonObj, _ := engine.LoadPrefab(res.Prefabs_button)
	//
	//	inputText := inputObj.FindComponentByName("TextView", true).(*builtin.TextView)
	//	objectText := object.FindComponentByName("TextView").(*builtin.TextView).Text
	//	engine.LogDebug(objectText)
	//	if objectText != "" {
	//		inputText.SetText(objectText)
	//	}
	//
	//	buttonObj.FindComponentByName("TextView", true).(*builtin.TextView).SetText("ok")
	//
	//	buttonObj.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
	//		newText := inputText.GetText()
	//		object.FindComponentByName("TextView", true).(*builtin.TextView).SetText(newText)
	//		engine.LogDebug("New Text: ", newText)
	//		return true
	//	}))
	//
	//	gridObject.AddChild(inputObj)
	//	gridObject.AddChild(buttonObj)
	//
	//	s.hierarchy.AddChild(gridObject)
	//}
	//
	//if object.GetName() == "Image Box" {
	//	engine.LogDebug("clicked on imageBox")
	//	gridObject := engine.NewSceneObject("image bruh")
	//
	//	grid := builtin.NewGridLayout()
	//	grid.Cols = 2
	//	grid.Rows = 1
	//
	//	gridObject.AddComponent(grid)
	//
	//	inputObj, _ := engine.LoadPrefab(res.Prefabs_inputBox)
	//	buttonObj, _ := engine.LoadPrefab(res.Prefabs_button)
	//
	//	inputText := inputObj.FindComponentByName("TextView", true).(*builtin.TextView)
	//	objectImageURL := object.FindComponentByName("ImageView").(*builtin.ImageView).ImageUrl
	//	engine.LogDebug(objectImageURL)
	//	if objectImageURL != "" {
	//		inputText.SetText(objectImageURL)
	//	}
	//	buttonObj.FindComponentByName("TextView", true).(*builtin.TextView).SetText("ok")
	//
	//	buttonObj.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
	//		url := "res/images/" + inputText.GetText()
	//		object.FindComponentByName("ImageView", true).(*builtin.ImageView).SetImageUrl(url)
	//		engine.LogDebug("New url: ", url)
	//		return true
	//	}))
	//
	//	gridObject.AddChild(inputObj)
	//	gridObject.AddChild(buttonObj)
	//
	//	s.hierarchy.AddChild(gridObject)
	//}
}
