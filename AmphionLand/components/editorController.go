package components

import (
	"AmphionLand/generated/res"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"strings"
)

type EditorController struct {
	engine.ComponentImpl
}

func (s *EditorController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	sceneObject1 := engine.NewSceneObject("left_scene")
	sceneObject1.Transform.Position.Z = 1
	sceneObject1.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,0)
	view := builtin.NewShapeView(builtin.ShapeRectangle)
	view.FillColor = a.NewColor(100, 100, 100)
	view.StrokeWeight = 0
	sceneObject1.AddComponent(view)

	sceneObject2 := engine.NewSceneObject("right_thing")
	sceneObject2.Transform.Position.Z = 1
	sceneObject2.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,0)
	layout := builtin.NewGridLayout()
	layout.Cols = 2
	layout.RowPadding = 20
	sceneObject2.AddComponent(layout)

	prefab, err := engine.LoadPrefab(res.Prefabs_button)

	if err==nil{
		prefab.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		textView := prefab.GetChildByName("Button text").GetComponentByName(".+TextView").(*builtin.TextView)
		textView.SetText("Save Prefab")
		textView.SetHTextAlign(a.TextAlignCenter)
		textView.SetVTextAlign(a.TextAlignCenter)
		sceneObject2.AddChild(prefab)
	}

	prefab2, err2 := engine.LoadPrefab(res.Prefabs_button)

	if err2==nil{
		prefab2.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		textView := prefab2.GetChildByName("Button text").GetComponentByName(".+TextView").(*builtin.TextView)
		textView.SetText("Save Scene")
		textView.SetHTextAlign(a.TextAlignCenter)
		textView.SetVTextAlign(a.TextAlignCenter)
		sceneObject2.AddChild(prefab2)
	}

	sceneObject2_3 := engine.NewSceneObject("Prefab List")
	sceneObject2_3.Transform.Position.Z = 1
	sceneObject2_3.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	view4 := builtin.NewShapeView(builtin.ShapeRectangle)
	view4.FillColor = a.NewColor(115, 115, 180)
	view4.StrokeWeight = 1
	sceneObject2_3.AddComponent(view4)
	prefabs_layout := builtin.NewGridLayout()
	prefabs_layout.RowPadding = 30
	sceneObject2_3.AddComponent(prefabs_layout)
	sceneObject2.AddChild(sceneObject2_3)
	s.SpawnPrefabsList(sceneObject2_3)

	sceneObject2_4 := engine.NewSceneObject("Hierarchy")
	sceneObject2_3.Transform.Position.Z = 1
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

func (s *EditorController) SpawnPrefabsList(sceneO *engine.SceneObject) {
	file, err := s.Engine.GetResourceManager().ReadFile(res.Strings_prefabsList)
	if err!=nil {
		engine.LogDebug("Failed to read file")
		engine.LogError(err.Error())
		return
	}
	fileStr := string(file)
	fileStrList := strings.Split(fileStr, "\r\n")
	for _, prefabName := range fileStrList {
		path := "prefabs/" + prefabName + ".yaml"
		prefab, err := engine.LoadPrefab(res.Prefabs_prefabViewer)
		if err!=nil {
			engine.LogError(err.Error())
			return
		}
		prefabViewerController := prefab.GetComponentByName(".+PrefabViewerController").(*PrefabViewerController)
		prefabViewerController.PrefabPath = path
		prefabViewerController.SetTextView(prefabName)

		sceneO.AddChild(prefab)
	}
}