package components

import (
	"AmphionLand/generated/res"

	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type TestComponent struct {
	engine.ComponentImpl
}

func (t *TestComponent) OnInit(ctx engine.InitContext) {
	t.ComponentImpl.OnInit(ctx)

	button, err := engine.LoadPrefab(res.Builtin_prefabs_button)
	if err != nil {
		engine.LogError(err.Error())
	}

	bg := builtin.NewShapeView(builtin.ShapeRectangle)
	bg.FillColor = a.GreenColor()
	bg.StrokeWeight = 0
	bg.CornerRadius = 10
	button.AddComponent(bg)

	onClickListener := button.GetComponentByName(".+OnClickListener").(*builtin.OnClickListener)
	onClickListener.OnClick = func(event engine.AmphionEvent) bool {
		engine.LogInfo("Prefab button clicked!")
		return true
	}

	t.SceneObject.AddChild(button)
}

func (t *TestComponent) GetName() string {
	return engine.NameOfComponent(t)
}
