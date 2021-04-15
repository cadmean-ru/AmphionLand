package components

import (
	"AmphionLand/generated/res"
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
	sceneObject1.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	view := builtin.NewShapeView(builtin.ShapeRectangle)
	view.FillColor = a.BlueColor()
	view.StrokeWeight = 0
	sceneObject1.AddComponent(view)

	sceneObject2 := engine.NewSceneObject("bruh2")
	sceneObject2.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	layout := builtin.NewGridLayout()
	layout.Cols = 2
	sceneObject2.AddComponent(layout)

	prefab, err := engine.LoadPrefab(res.Prefabs_button)

	if err==nil{
		prefab.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		view2 := builtin.NewShapeView(builtin.ShapeRectangle)
		view2.FillColor = a.BlueColor()
		view2.CornerRadius = 69
		prefab.AddComponent(view2)
		sceneObject2.AddChild(prefab)
	}

	prefab2, err2 := engine.LoadPrefab(res.Prefabs_button)

	if err2==nil{
		prefab2.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		view3 := builtin.NewShapeView(builtin.ShapeRectangle)
		view3.FillColor = a.GreenColor()
		view3.CornerRadius = 69
		prefab2.AddComponent(view3)
		sceneObject2.AddChild(prefab2)
	}

	s.SceneObject.AddChild(sceneObject1)
	s.SceneObject.AddChild(sceneObject2)
}

func (s *EditorController) OnStart() {

}

func (s *EditorController) GetName() string {
	return engine.NameOfComponent(s)
}