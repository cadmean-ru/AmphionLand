package components

import (
	"AmphionLand/generated/res"
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

	//radioButt, err := engine.LoadPrefab(res.Prefabs_radioButtonPrefab)
	//if err == nil {
	//	l.SceneObject.AddChild(radioButt)
	//	engine.LogDebug("Here")
	//} else {
	//	engine.LogDebug(err.Error())
	//}

	obj := engine.NewSceneObject("Radio butts")
	obj.SetSizeXy(100, 100)

	radioButt := NewRadioButtonGroup()
	radioButt.AddItem("test 1")
	radioButt.AddItem("test 2")
	radioButt.AddItem("test 3")

	obj.AddComponent(radioButt)

	l.SceneObject.AddChild(obj)

	butt, _ := engine.LoadPrefab(res.Builtin_prefabs_button)
	butt.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
		engine.LogDebug(radioButt.SelectedItemText())
		return true
	}))
	butt.Transform.Position.X = 200
	l.SceneObject.AddChild(butt)

	l.paddingObject = engine.NewSceneObject("padding")
	l.paddingObject.AddComponent(NewPadding())
	l.paddingObject.AddComponent(builtin.NewBoundaryView())
	butt.AddComponent(builtin.NewBoundaryView())
	l.SceneObject.AddChild(l.paddingObject)

	butt2, _ := engine.LoadPrefab(res.Builtin_prefabs_button)
	butt2.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event2 engine.AmphionEvent) bool {
		engine.LogDebug("butt 2")
		return true
	}))
	butt2.FindComponentByName("TextView", true).(*builtin.TextView).SetText("butt2")
	l.paddingObject.AddChild(butt2)


	//
	//wodgetPrefab, err := engine.LoadPrefab(res.Prefabs_wodget)
	//if err == nil {
	//	l.SceneObject.AddChild(wodgetPrefab)
	//	engine.LogDebug("Here")
	//} else {
	//	engine.LogDebug(err.Error())
	//}

	//wodget := engine.NewSceneObject("wodget")
	//wodget.Transform.Size = a.NewVector3(68, 68, 0)
	//wodget.Transform.Position = a.NewVector3(69, 69, 69)
	//wodget.AddComponent(builtin.NewRectBoundary())
	//rect := builtin.NewShapeView(builtin.ShapeRectangle)
	//rect.FillColor = a.NewColor("#2c69a8")
	//wodget.AddComponent(rect)
	//wodget.AddComponent(builtin.NewMouseMover())
	//l.SceneObject.AddChild(wodget)

	engine.LogDebug("OnInit")
}

func (l *LoginSceneController) OnStart() {
	engine.LogDebug("OnStart 2")
	paddingComponent := l.paddingObject.FindComponentByName("Padding", true).(*Padding)
	paddingComponent.LeftX = 100
	paddingComponent.UpY = 100
	paddingComponent.DownY = 50
	paddingComponent.UpdatePadding()
}

func (l *LoginSceneController) GetName() string {
	return engine.NameOfComponent(l)
}
