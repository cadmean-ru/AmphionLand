package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)


type WidgetGrid struct {
	builtin.GridLayout
	Size int `state`
}

func (s *WidgetGrid) OnInit(ctx engine.InitContext) {
	s.GridLayout.OnInit(ctx)
	s.Size = 5
	for i := 0; i < s.Size; i++ {
		s.AddColumn(a.WrapContent)
		s.AddRow(a.WrapContent)
	}

	//engine.LogDebug("Here %+v\n\n\n", s.Size)

	for i := 0; i < s.Size * s.Size; i++ {
		//engine.LogDebug("Here %+v\n\n\n", s.SceneObject)
		sceneObj := engine.NewSceneObject(fmt.Sprint("go barbruh go ", i))
		sceneObj.Transform.Position = a.NewVector3(0, 0, 0)
		sceneObj.Transform.Size.X = 10
		sceneObj.Transform.Size.Y = 100
		sceneObj.AddComponent(&EmptyBox{})
		sceneObj.AddComponent(builtin.NewRectBoundary())
		ell := builtin.NewShapeView(builtin.ShapeEllipse)
		ell.FillColor = a.NewColor(255, 255, 0)
		sceneObj.AddComponent(ell)
		s.SceneObject.AddChild(sceneObj)
	}

	//engine.RunTask(engine.NewTaskBuilder().Run(func() (interface{}, error) {
	//	time.Sleep(100)
	//	s.SceneObject.ForEachObject(func(object *engine.SceneObject) {
	//		engine.LogDebug(object.GetName())
	//		engine.LogDebug("%+v", object.Transform.Position)
	//	})
	//	return nil, nil
	//}).Build())
}

func (s *WidgetGrid) GetName() string {
	return engine.NameOfComponent(s)
}
