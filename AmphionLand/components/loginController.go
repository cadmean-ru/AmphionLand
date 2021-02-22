package components

import (
	"AmphionLand/generated/res"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"github.com/cadmean-ru/amphion/engine/rpc"
)

type LoginSceneController struct {
	engine.ComponentImpl
	emailInput *builtin.NativeInputView
	passwordInput *builtin.NativeInputView
}

func (l *LoginSceneController) OnInit(ctx engine.InitContext) {
	l.ComponentImpl.OnInit(ctx)

	l.emailInput = l.SceneObject.GetChildByName("email input").GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.NativeInputView").(*builtin.NativeInputView)
	l.passwordInput = l.SceneObject.GetChildByName("password input").GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.NativeInputView").(*builtin.NativeInputView)

	buttonPrefab, err := engine.LoadPrefab(res.Builtin_prefabs_button)
	if err == nil {
		textObj := buttonPrefab.GetChildByName("Button text")
		textView := textObj.GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView)
		textView.Text = "login"
		textView.TextColor = a.NewColor("#fff")

		clickListener := buttonPrefab.GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.OnClickListener").(*builtin.OnClickListener)
		clickListener.OnClick = func(event engine.AmphionEvent) bool {
			email := l.emailInput.GetText()
			password := l.passwordInput.GetText()

			rpc.F("login").Then(func(res interface{}) {
				engine.LogDebug("%+v", res)
			}).Err(func(err error) {
				engine.LogDebug("%+v", err)
			}).Call(email, password)

			return true
		}

		l.SceneObject.AddChild(buttonPrefab)
	}
}

func (l *LoginSceneController) GetName() string {
	return engine.NameOfComponent(l)
}
