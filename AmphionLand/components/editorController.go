package components

import (
	"AmphionLand/generated/res"
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type EditorController struct {
	engine.ComponentImpl
	yeetingSceneObject *engine.SceneObject
	hierarchy *engine.SceneObject
	remover bool
	containerPrefab *engine.SceneObject
	showBoxed bool
}

func (s *EditorController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	leftScene := engine.NewSceneObject("left_scene")
	leftScene.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,0)
	view := builtin.NewShapeView(builtin.ShapeRectangle)
	view.FillColor = a.NewColor(100, 100, 100)
	view.StrokeWeight = 0
	leftScene.AddComponent(view)
	layout := builtin.NewGridLayout()
	layout.AddColumn(a.FillParent)
	leftScene.AddComponent(layout)

	engine.BindEventHandler(engine.EventKeyDown, func(event engine.AmphionEvent) bool {
		engine.LogDebug("key pressed: %s", event.StringData())
		return true
	})

	engine.BindEventHandler(engine.EventKeyUp, func(event engine.AmphionEvent) bool {
		engine.LogDebug("key released: %s", event.StringData())
		return true
	})

	engine.BindEventHandler(engine.EventTextInput, func(event engine.AmphionEvent) bool {
		engine.LogDebug("text input: %s", string(event.StringData()))
		return true
	})

	engine.GetCurrentScene().AddComponent(NewNewNewScrollManager())

	//s.containerPrefab, _ = engine.LoadPrefab(res.Prefabs_editorContainer)
	//engine.LogDebug("bruhsdfsd")
	//for _, c := range s.containerPrefab.GetComponents(true) {
	//	engine.LogDebug(c.GetName())
	//}

	for i := 0; i < 15; i++ {
		box := engine.NewSceneObject(fmt.Sprintf("Box %d", i))
		rectangle := builtin.NewShapeView(builtin.ShapeRectangle)
		rectangle.FillColor = a.TransparentColor()
		rectangle.StrokeColor = a.BlackColor()
		rectangle.StrokeWeight = 0
		box.Transform.Size.Y = 100
		box.AddComponent(rectangle)
		box.AddComponent(builtin.NewRectBoundary())
		box.AddComponent(&ClickAndInspeceet{})
		leftScene.AddChild(box)
	}

	rightThing := engine.NewSceneObject("right_thing")
	rightThing.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,0)
	layoutRight := builtin.NewGridLayout()
	layoutRight.AddColumn(a.FillParent)
	layoutRight.AddColumn(a.FillParent)
	layoutRight.RowPadding = 3
	rightThing.AddComponent(layoutRight)

	playButton, err := engine.LoadPrefab(res.Prefabs_button)

	if err==nil{
		playButton.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		textView := playButton.FindComponentByName("TextView", true).(*builtin.TextView)
		textView.SetText("Test app")
		textView.SetHTextAlign(a.TextAlignCenter)
		textView.SetVTextAlign(a.TextAlignCenter)
		playButton.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
			engine.CloseScene(func() {
				newScene := leftScene.Copy(leftScene.GetName())
				newScene.ForEachObject(func(object *engine.SceneObject) {
					object.RemoveComponentByName("ClickAndInspeceet")
					object.RemoveComponentByName("Yeeter")
					object.RemoveComponentByName("EditorGrid")
				}, true)
				_= engine.ShowScene(newScene)
			})
			return true
		}))

		rightThing.AddChild(playButton)
	}

	prefab2, err2 := engine.LoadPrefab(res.Prefabs_button)

	if err2==nil{
		prefab2.Transform.Size = a.NewVector3(a.MatchParent,50,4)
		prefab2.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
			savedScene := leftScene.Copy("savedScene")
			savedScene.ForEachObject(func(object *engine.SceneObject) {
				if len(object.GetComponentsByName("ClickAndInspeceet", true)) > 0 {
					object.RemoveAllComponents()
				}
			}, true)

			yaml, err := savedScene.EncodeToYaml()
			if err != nil {
				return false
			}

			err = ioutil.WriteFile("./savedScene.yaml", yaml, 0644)
			if err != nil{
				engine.LogDebug(err.Error())
			}

			return true
		}))
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

		gridViewer.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
			s.showBoxed = !s.showBoxed

			leftScene.ForEachObject(func(object *engine.SceneObject) {
				comps := object.GetComponentsByName("ClickAndInspeceet")
				if len(comps) == 0 {
					return
				}

				comp := comps[0].(*ClickAndInspeceet)
				comp.ToggleBox()
			})
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
			prefabRemoveShapeView.Redraw()
			engine.RequestRendering()
			return true
		}))
		rightThing.AddChild(prefabRemove)
	}

	sceneobject23 := engine.NewSceneObject("Prefab List")
	sceneobject23.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	view4 := builtin.NewShapeView(builtin.ShapeRectangle)
	view4.FillColor = a.NewColor(115, 115, 180)
	view4.StrokeWeight = 1
	sceneobject23.AddComponent(view4)
	prefabsLayout := builtin.NewGridLayout()
	prefabsLayout.AddColumn(a.FillParent)
	prefabsLayout.RowPadding = 0
	sceneobject23.AddComponent(prefabsLayout)
	rightThing.AddChild(sceneobject23)

	s.SpawnPrefabsList(sceneobject23)

	s.hierarchy = engine.NewSceneObject("Hierarchy")
	s.hierarchy.Transform.Size = a.NewVector3(a.MatchParent,a.MatchParent,1)
	view5 := builtin.NewShapeView(builtin.ShapeRectangle)
	view5.FillColor = a.NewColor(115, 115, 180)
	view5.StrokeWeight = 1
	s.hierarchy.AddComponent(view5)
	hierarchyLayout := builtin.NewGridLayout()
	hierarchyLayout.AddColumn(a.FillParent)
	s.hierarchy.AddComponent(hierarchyLayout)
	rightThing.AddChild(s.hierarchy)

	s.SceneObject.AddChild(leftScene)
	s.SceneObject.AddChild(rightThing)
}

func (s *EditorController) OnStart() {
	apiKey, err := s.Engine.GetResourceManager().ReadFile(res.Strings_definetlynotkey)
	if err==nil {
		w, apiErr := owm.NewCurrent("C", "ru", string(apiKey))
		if apiErr == nil {
			_ = w.CurrentByName("Moscow")
			engine.LogDebug(strconv.FormatFloat(w.Main.Temp, 'f',3,32))
			engine.LogDebug(strconv.FormatFloat(w.Main.FeelsLike, 'f',3,32))
			//engine.LogDebug(strconv.FormatFloat(w.Main.GrndLevel, 'f',3,32))
			engine.LogDebug(strconv.FormatFloat(w.Main.TempMax, 'f',3,32))
			engine.LogDebug(strconv.FormatFloat(w.Main.TempMin, 'f',3,32))
			//engine.LogDebug(string(rune(w.Main.Humidity)))
			engine.LogDebug(w.Weather[0].Description)
		} else {
			engine.LogDebug(apiErr.Error())
		}
	} else {
		engine.LogDebug(err.Error())
	}

	engine.RunTask(engine.NewTaskBuilder().Run(func() (interface{}, error) {
		time.Sleep(1000)
		engine.FindObjectByName("left_scene").ForEachChild(func(box *engine.SceneObject) {
			engine.LogDebug("min y %f", box.Transform.GetGlobalRect().GetMin().Y)
		})
		return nil, nil
	}).Build())
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