package components

import (
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type LoginSceneController struct {
	engine.ComponentImpl
	emailInput *builtin.NativeInputView
	passwordInput *builtin.NativeInputView
	paddingObject *engine.SceneObject
}

func (l *LoginSceneController) OnInit(ctx engine.InitContext) {
	l.ComponentImpl.OnInit(ctx)

	//obj := engine.NewSceneObject("Radio butts")
	//obj.SetSizeXy(100, 100)
	//obj.SetPositionXy(10, 10)
	//
	//radioButt := NewRadioButtonGroup()
	//radioButt.AddItem("test 1")
	//radioButt.AddItem("test 2")
	//radioButt.AddItem("test 3")
	//
	//radioButt.SetOnItemSelectedListener(func(item RadioItem) {
	//	engine.LogDebug("Selected '%s'", item.text)
	//})
	//
	//radioButt.SetSelected(0)
	//
	//obj.AddComponent(radioButt)
	//
	//l.SceneObject.AddChild(obj)
	//
	//butt, _ := engine.LoadPrefab(res.Builtin_prefabs_button)
	//butt.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
	//	engine.LogDebug(radioButt.SelectedItemText())
	//	return true
	//}))
	//butt.Transform.Position.X = 200
	//l.SceneObject.AddChild(butt)

	checkBoxObj := engine.NewSceneObject("Check butts")
	checkBoxObj.SetSizeXy(100, 100)

	checkButt := NewCheckBoxGroup()
	checkButt.AddItem("test 1")
	checkButt.AddItem("test 2")
	checkButt.AddItem("test 3")

	checkBoxObj.AddComponent(checkButt)

	l.SceneObject.AddChild(checkBoxObj)

	//butt, _ := engine.LoadPrefab(res.Builtin_prefabs_button)
	//butt.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
	//	indexes := ""
	//	for i := 0; i < len(checkButt.selectedIndexes); i++ {
	//		indexes += strconv.Itoa(checkButt.selectedIndexes[i]) + " "
	//	}
	//	engine.LogDebug(indexes)
	//	return true
	//}))
	//butt.Transform.Position.X = 200
	//l.SceneObject.AddChild(butt)
	//
	//bigGridObj := engine.NewSceneObject("bigGrid")
	//
	//bigGrid := builtin.NewGridLayoutSized(1, 2)
	//
	//bigGridObj.AddComponent(bigGrid)
	////bigGrid.Cols = 3
	////bigGrid.Rows = 3
	////bigGrid.RowPadding = 10
	////bigGrid.ColPadding = 10
	//
	//l.SceneObject.AddChild(bigGridObj)
	//
	//butt3, _ := engine.LoadPrefab(res.Builtin_prefabs_button)
	//butt3.FindComponentByName("TextView", true).(*builtin.TextView).SetText("butt3")
	////bigGridObj.AddChild(butt3)
	//
	//paddingObject2 := engine.NewSceneObject("padding2")
	//paddingObject2.Transform.Size = a.NewVector3(100, 100, 1)
	//paddingObject2.AddComponent(NewPadding())
	//paddingObject2.AddComponent(builtin.NewBoundaryView())
	//paddingObject2.AddChild(butt3)
	//bigGridObj.AddChild(paddingObject2)
	//
	//butt2, _ := engine.LoadPrefab(res.Builtin_prefabs_button)
	////butt2.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event2 engine.AmphionEvent) bool {
	////	engine.LogDebug("butt 2")
	////	return true
	////}))
	//butt2.FindComponentByName("TextView", true).(*builtin.TextView).SetText("butt2")
	////bigGridObj.AddChild(butt2)
	//l.paddingObject = engine.NewSceneObject("padding")
	//l.paddingObject.Transform.Size = a.NewVector3(100, 100, 1)
	//l.paddingObject.AddComponent(NewPadding())
	//l.paddingObject.AddComponent(builtin.NewBoundaryView())
	//l.paddingObject.AddChild(butt2)
	//bigGridObj.AddChild(l.paddingObject)
	//
	//
	////
	////wodgetPrefab, err := engine.LoadPrefab(res.Prefabs_wodget)
	////if err == nil {
	////	l.SceneObject.AddChild(wodgetPrefab)
	////	engine.LogDebug("Here")
	////} else {
	////	engine.LogDebug(err.Error())
	////}
	//
	////wodget := engine.NewSceneObject("wodget")
	////wodget.Transform.Size = a.NewVector3(68, 68, 0)
	////wodget.Transform.Position = a.NewVector3(69, 69, 69)
	////wodget.AddComponent(builtin.NewRectBoundary())
	////rect := builtin.NewShapeView(builtin.ShapeRectangle)
	////rect.FillColor = a.NewColor("#2c69a8")
	////wodget.AddComponent(rect)
	////wodget.AddComponent(builtin.NewMouseMover())
	////l.SceneObject.AddChild(wodget)

	engine.LogDebug("OnInit")
}

func (l *LoginSceneController) OnStart() {
	//engine.LogDebug("OnStart 2")
	//paddingComponent := l.paddingObject.FindComponentByName("Padding", true).(*Padding)
	//paddingComponent.LeftX = 100
	//paddingComponent.UpY = 100
	//paddingComponent.DownY = 50
	//paddingComponent.RightX = 25
	//paddingComponent.UpdatePadding()
}

func (l *LoginSceneController) GetName() string {
	return engine.NameOfComponent(l)
}
