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
}

func (l *LoginSceneController) OnInit(ctx engine.InitContext) {
	l.ComponentImpl.OnInit(ctx)

	wodgetPrefab, err := engine.LoadPrefab(res.Prefabs_wodget)
	if err == nil {
		l.SceneObject.AddChild(wodgetPrefab)
	}
}

func (l *LoginSceneController) GetName() string {
	return engine.NameOfComponent(l)
}
