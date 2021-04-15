package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type EditorController struct {
	engine.ComponentImpl
}

func (s *EditorController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	sceneObject1 := engine.NewSceneObject("bruh1")
	sceneObject1.Transform.Size = a.NewVector3(50,50,50)
	view := builtin.NewShapeView(builtin.ShapeRectangle)
	view.FillColor = a.BlueColor()
	sceneObject1.AddComponent(view)

	s.SceneObject.AddChild(sceneObject1)
}

func (s *EditorController) OnStart() {

}

func (s *EditorController) GetName() string {
	return engine.NameOfComponent(s)
}