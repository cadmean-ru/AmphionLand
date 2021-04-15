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
	view.FillColor = a.NewColor(100, 100, 100)
	view.StrokeWeight = 0
	sceneObject1.AddComponent(view)

	sceneObject2 := engine.NewSceneObject("bruh2")
	sceneObject2.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	layout := builtin.NewGridLayout()
	layout.Cols = 2
	layout.RowPadding = 20
	sceneObject2.AddComponent(layout)

	prefab, err := engine.LoadPrefab(res.Prefabs_button)

	if err==nil{
		prefab.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		view2 := builtin.NewShapeView(builtin.ShapeRectangle)
		view2.FillColor = a.NewColor(102, 102, 153)
		view2.CornerRadius = 69
		textView := prefab.GetChildByName("Button text").GetComponentByName(".+TextView").(*builtin.TextView)
		textView.SetText("Save Prefab")
		textView.SetHTextAlign(a.TextAlignCenter)
		textView.SetVTextAlign(a.TextAlignCenter)
		prefab.AddComponent(view2)
		sceneObject2.AddChild(prefab)
	}

	prefab2, err2 := engine.LoadPrefab(res.Prefabs_button)

	if err2==nil{

		prefab2.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		view3 := builtin.NewShapeView(builtin.ShapeRectangle)
		view3.FillColor = a.NewColor(102, 102, 153)
		view3.CornerRadius = 69
		textView := prefab2.GetChildByName("Button text").GetComponentByName(".+TextView").(*builtin.TextView)
		textView.SetText("Save Scene")
		textView.SetHTextAlign(a.TextAlignCenter)
		textView.SetVTextAlign(a.TextAlignCenter)
		prefab2.AddComponent(view3)
		sceneObject2.AddChild(prefab2)
	}

	sceneObject2_3 := engine.NewSceneObject("bruh3")
	sceneObject2_3.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	view4 := builtin.NewShapeView(builtin.ShapeRectangle)
	view4.FillColor = a.NewColor(115, 115, 180)
	view4.StrokeWeight = 1
	sceneObject2_3.AddComponent(view4)
	sceneObject2.AddChild(sceneObject2_3)

	sceneObject2_4 := engine.NewSceneObject("bruh4")
	sceneObject2_4.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	view5 := builtin.NewShapeView(builtin.ShapeRectangle)
	view5.FillColor = a.NewColor(115, 115, 180)
	view5.StrokeWeight = 1
	sceneObject2_4.AddComponent(view5)
	sceneObject2.AddChild(sceneObject2_4)

	s.SceneObject.AddChild(sceneObject1)
	s.SceneObject.AddChild(sceneObject2)
}

func (s *EditorController) OnStart() {

}

func (s *EditorController) GetName() string {
	return engine.NameOfComponent(s)
}