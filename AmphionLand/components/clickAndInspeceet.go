package components

import (
	"AmphionLand/generated/res"
	"fmt"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/common/require"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"strconv"
)

type ClickAndInspeceet struct {
	engine.ComponentImpl
	editor    *EditorController
	hierarchy *engine.SceneObject
	shape     *builtin.ShapeView
	cm        *engine.ComponentsManager
}

func (s *ClickAndInspeceet) OnStart() {
	engine.LogDebug("inspeceet start")
	engine.BindEventHandler(engine.EventMouseDown, s.handleClick)

	s.editor = FindEditorController(engine.GetCurrentScene())
	s.hierarchy = engine.FindObjectByName("Hierarchy")
	s.shape = builtin.GetShapeView(s.SceneObject)

	s.cm = engine.GetInstance().GetComponentsManager()
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
		editorGrid := GetEditorGrid(newObj)
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

	transMap := map[string]a.Vector3{
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
		grid.AddColumn(a.FillParent)
		grid.AddColumn(a.FillParent)
		grid.AddRow(a.FillParent)
		grid.AddRow(a.FillParent)
		grid.AddRow(a.FillParent)

		box.AddComponent(grid)

		nameMap := [3]string{"X", "Y", "Z"}
		valueMap := [3]float32{vec.X, vec.Y, vec.Z}

		for j := 0; j < 3; j++ {
			name := engine.NewSceneObject("name" + string(rune(i)) + string(rune(j)))
			nameText := builtin.NewTextView(nameMap[j])
			nameText.SetVTextAlign(a.TextAlignCenter)
			nameText.SetHTextAlign(a.TextAlignCenter)

			name.AddComponent(nameText)
			name.Transform.Size.Y = 30

			value := engine.NewSceneObject("value" + string(rune(i)) + string(rune(j)))
			valueText := builtin.NewTextView(strconv.FormatFloat(float64(valueMap[j]), 'f', 3, 32))
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

	for _, comp := range components {

		componentsSomething := engine.NewSceneObject("King of components " + engine.NameOfComponent(comp))

		publics := s.cm.GetComponentState(comp)
		componentsSomething.Transform.Size.Y = float32(30 * (len(publics) + 1))
		componentsSomething.AddComponent(builtin.NewBoundaryView())
		grid := builtin.NewGridLayout()
		grid.AddColumn(a.FillParent)
		grid.AddColumn(a.FillParent)

		componentsSomething.AddComponent(grid)

		nameBoxLabel := engine.NewSceneObject("nameBoxLabel")
		nameBoxLabel.Transform.Size.Y = 30
		nameBoxLabelText := builtin.NewTextView("Name: ")
		nameBoxLabelText.SetVTextAlign(a.TextAlignCenter)
		nameBoxLabelText.SetHTextAlign(a.TextAlignLeft)
		nameBoxLabel.AddComponent(nameBoxLabelText)
		componentsSomething.AddChild(nameBoxLabel)

		nameBox := engine.NewSceneObject("nameBox " + engine.NameOfComponent(comp))
		nameBox.Transform.Size.Y = 30
		nameBoxText := builtin.NewTextView(engine.NameOfComponent(comp))
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

			switch public.(type) {
			case string, int, int32, int64, uint, uint8, float32, float64:
				s.CreateInputBox(public, name, comp, componentsSomething)
			case bool:
				s.CreateCheckBox(public, name, comp, componentsSomething)
			default:
				s.CreateInputBox(public, name, comp, componentsSomething)
			}


			engine.LogDebug("pub %s %v", name, public)
		}
	}
}

func (s *ClickAndInspeceet) CreateCheckBox(public interface{}, name string, comp engine.Component,
componentsSomething *engine.SceneObject) {
	stateCheck, _ := engine.LoadPrefab(res.Prefabs_checkBox)
	checkBoxGroup := stateCheck.FindComponentByName("CheckBoxGroup", true).(*CheckBoxGroup)
	checkBoxGroup.AddItem("")
	if public == true {
		checkBoxGroup.SetSelected(0)
	}

	nameInput := name
	compus := comp
	checkBoxGroup.SetOnItemSelectedListener(func(item CheckItem) {
		compusState := s.cm.GetComponentState(compus)
		if compusState[nameInput] == true {
			compusState[nameInput] = false
		} else {
			compusState[nameInput] = true
		}
		s.cm.SetComponentState(compus, compusState)

		engine.ForceAllViewsRedraw()
		engine.RequestRendering()
	})

	engine.ForceAllViewsRedraw()
	engine.RequestRendering()
	componentsSomething.AddChild(stateCheck)
}


func (s *ClickAndInspeceet) CreateInputBox(public interface{}, name string, comp engine.Component,
componentsSomething *engine.SceneObject) {
	stateInput, _ := engine.LoadPrefab(res.Prefabs_inputBox)
	stateInput.Transform.Size.Y = 30
	inputField := stateInput.FindComponentByName("InputField", true).(*InputField)
	inputField.allowParagraph = false
	switch public.(type) {
	case string:
		inputField.varType = stringType
	case int, int32, int64, uint, uint8:
		inputField.varType = intType
	case float32, float64:
		inputField.varType = floatType
	}
	inputField.SetText(require.String(public))

	typeInput := public
	nameInput := name
	compus := comp
	inputField.someAction = func() {
		var newValue interface{}
		switch typeInput.(type) {
		case string:
			newValue = require.String(inputField.text)
		case int, int32, int64, uint, uint8:
			newValue = require.Int(inputField.text)
		case a.TextAlign:
			newValue = a.TextAlign(require.Byte(inputField.text))
		case float32:
			newValue = require.Float32(inputField.text)
		case float64:
			newValue = require.Float64(inputField.text)
		}
		compusState := s.cm.GetComponentState(compus)
		compusState[nameInput] = newValue
		s.cm.SetComponentState(compus, compusState)

		engine.ForceAllViewsRedraw()
		engine.RequestRendering()
	}
	engine.ForceAllViewsRedraw()
	engine.RequestRendering()
	componentsSomething.AddChild(stateInput)
}
