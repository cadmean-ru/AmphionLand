package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type LoginSceneController struct {
	engine.ComponentImpl
	emailInput *builtin.NativeInputView
	passwordInput *builtin.NativeInputView
}

func (l *LoginSceneController) OnInit(ctx engine.InitContext) {
	l.ComponentImpl.OnInit(ctx)

	//wodgetPrefab, err := engine.LoadPrefab(res.Prefabs_wodget)
	//if err == nil {
	//	l.SceneObject.AddChild(wodgetPrefab)
	//}

	wodget := engine.NewSceneObject("wodget")
	wodget.Transform.Size = a.NewVector3(68, 68, 0)
	wodget.Transform.Position = a.NewVector3(69, 69, 69)
	wodget.AddComponent(builtin.NewRectBoundary())
	wodget.AddComponent(builtin.NewShapeView(builtin.ShapeRectangle))
	wodget.AddComponent(builtin.NewMouseMover())
	l.SceneObject.AddChild(wodget)
}

func (l *LoginSceneController) GetName() string {
	return engine.NameOfComponent(l)
}
