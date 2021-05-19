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
	yeetingSceneObject *engine.SceneObject
	hierarchy *engine.SceneObject
	remover bool
}

func (s *EditorController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	leftScene := engine.NewSceneObject("left_scene")
	leftScene.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,0)
	view := builtin.NewShapeView(builtin.ShapeRectangle)
	view.FillColor = a.NewColor(100, 100, 100)
	view.StrokeWeight = 0
	leftScene.AddComponent(view)
	leftScene.AddComponent(builtin.NewGridLayout())
	for i := 0; i < 10; i++ {
		box := engine.NewSceneObject("emptyBox" + string(rune(i)))
		rectangle := builtin.NewShapeView(builtin.ShapeRectangle)
		rectangle.FillColor = a.TransparentColor()
		rectangle.StrokeColor = a.BlackColor()
		rectangle.StrokeWeight = 0
		box.Transform.Size.Y = 100
		box.AddComponent(rectangle)
		box.AddComponent(&EmptyBox{})
		box.AddComponent(builtin.NewRectBoundary())
		box.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
			engine.LogDebug("ldskjfklsj")

			if s.yeetingSceneObject == nil {
				if s.remover {
					box.RemoveAllChildren()
					s.hierarchy.RemoveAllChildren()
				} else {

				}
				return false
			}

			engine.LogDebug("ldskjfklsj")

			box.RemoveAllChildren()

			s.yeetingSceneObject.RemoveComponentByName("Yeeter")
			s.yeetingSceneObject.SetParent(box)
			s.yeetingSceneObject = nil
			engine.LogDebug("here 1lekj")
			return false
		}))
		leftScene.AddChild(box)
	}

	rightThing := engine.NewSceneObject("right_thing")
	rightThing.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,0)
	layout := builtin.NewGridLayout()
	layout.Cols = 2
	layout.RowPadding = 20
	rightThing.AddComponent(layout)

	prefab, err := engine.LoadPrefab(res.Prefabs_button)

	if err==nil{
		prefab.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		textView := prefab.FindComponentByName("TextView", true).(*builtin.TextView)
		textView.SetText("Save Prefab")
		textView.SetHTextAlign(a.TextAlignCenter)
		textView.SetVTextAlign(a.TextAlignCenter)
		rightThing.AddChild(prefab)
	}

	prefab2, err2 := engine.LoadPrefab(res.Prefabs_button)

	if err2==nil{
		prefab2.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		textView := prefab2.FindComponentByName("TextView", true).(*builtin.TextView)
		textView.SetText("Save Scene")
		textView.SetHTextAlign(a.TextAlignCenter)
		textView.SetVTextAlign(a.TextAlignCenter)
		rightThing.AddChild(prefab2)
	}

	gridViewer, gridErr := engine.LoadPrefab(res.Prefabs_button)

	if gridErr==nil{
		gridViewer.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		textView := gridViewer.FindComponentByName("TextView", true).(*builtin.TextView)
		textView.SetText("Show Grid")
		textView.SetHTextAlign(a.TextAlignCenter)
		textView.SetVTextAlign(a.TextAlignCenter)

		flag := false
		gridViewer.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
			strokeWeight := 0
			if flag {
				strokeWeight = 0
				flag = !flag
			} else {
				strokeWeight = 1
				flag = !flag
			}
			for _, box := range leftScene.GetChildren(){
				rect := box.GetComponentByName("ShapeView").(*builtin.ShapeView)
				rect.StrokeWeight = byte(strokeWeight)
				rect.ForceRedraw()
			}
			engine.RequestRendering()
			return true
		}))
		rightThing.AddChild(gridViewer)
	}

	prefabRemove, prefabRemoveErr := engine.LoadPrefab(res.Prefabs_button)

	if prefabRemoveErr==nil{
		prefabRemove.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		textView := prefabRemove.FindComponentByName("TextView", true).(*builtin.TextView)
		textView.SetText("Remove Prefab")
		textView.SetHTextAlign(a.TextAlignCenter)
		textView.SetVTextAlign(a.TextAlignCenter)

		prefabRemoveShapeView := prefabRemove.GetComponentByName("ShapeView", true).(*builtin.ShapeView)
		prefabRemoveShapeViewColor := prefabRemoveShapeView.FillColor
		s.remover = false
		prefabRemove.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
			if s.remover {
				prefabRemoveShapeView.FillColor = prefabRemoveShapeViewColor
			} else {
				prefabRemoveShapeView.FillColor = a.NewColor(50,50,50)
			}
			s.remover = !s.remover
			prefabRemoveShapeView.ForceRedraw()
			engine.RequestRendering()
			return true
		}))
		rightThing.AddChild(prefabRemove)
	}

	sceneObject2_3 := engine.NewSceneObject("Prefab List")
	sceneObject2_3.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	view4 := builtin.NewShapeView(builtin.ShapeRectangle)
	view4.FillColor = a.NewColor(115, 115, 180)
	view4.StrokeWeight = 1
	sceneObject2_3.AddComponent(view4)
	prefabs_layout := builtin.NewGridLayout()
	prefabs_layout.RowPadding = 10
	sceneObject2_3.AddComponent(prefabs_layout)
	rightThing.AddChild(sceneObject2_3)

	s.SpawnPrefabsList(sceneObject2_3)

	s.hierarchy = engine.NewSceneObject("Hierarchy")
	s.hierarchy.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	view5 := builtin.NewShapeView(builtin.ShapeRectangle)
	view5.FillColor = a.NewColor(115, 115, 180)
	view5.StrokeWeight = 1
	s.hierarchy.AddComponent(view5)
	s.hierarchy.AddComponent(builtin.NewGridLayout())
	rightThing.AddChild(s.hierarchy)

	s.SceneObject.AddChild(leftScene)
	s.SceneObject.AddChild(rightThing)
}

func (s *EditorController) OnStart() {

}

func (s *EditorController) GetName() string {
	return engine.NameOfComponent(s)
}

func (s *EditorController) SpawnPrefabsList(sceneO *engine.SceneObject) {
	engine.RunTask(engine.NewTaskBuilder().Run(func() (interface{}, error) {
		return s.Engine.GetResourceManager().ReadFile(res.Strings_prefabsList)
	}).Then(func(res2 interface{}) {
		file := res2.([]byte)

		fileStr := string(file)
		fileStrList := strings.Split(fileStr, "\n")
		for _, prefabName := range fileStrList {
			prefabName = strings.ReplaceAll(prefabName, "\r", "")
			path := "prefabs/" + prefabName + ".yaml"
			prefab, err := engine.LoadPrefab(res.Prefabs_prefabViewer)
			if err!=nil {
				engine.LogError(err.Error())
				return
			}
			prefabViewerController := prefab.GetComponentByName("PrefabViewerController", true).(*PrefabViewerController)
			prefabViewerController.PrefabPath = path
			prefabViewerController.Text = prefabName

			sceneO.AddChild(prefab)
		}
	}).Err(func(err error) {
		engine.LogDebug("Failed to read file")
		engine.LogError(err.Error())
	}).Build())
}