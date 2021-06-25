package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type Selection struct{
	engine.ComponentImpl
	color a.Color
	componentData *builtin.TextView
	flag bool
}

func (s *Selection) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	s.Engine.BindEventHandler(engine.EventMouseDown, func(selectionEvent engine.AmphionEvent) bool {
		//engine.LogDebug(fmt.Sprintf("data: %+v", selectionEvent.Data.(engine.MouseEventData).SceneObject.GetName()))

		sceneObject := selectionEvent.Data.(engine.MouseEventData).SceneObject

		if sceneObject != nil {

			sceneObject.ForEachComponent(func(component engine.Component) {
				if view, ok := component.(*builtin.TextView); ok {
					engine.LogDebug(fmt.Sprintf("text: %+v", view.Text))
					s.componentData = view

					s.color = view.TextColor
					view.TextColor = a.RedColor()

					s.flag = true
					view.Redraw()
					engine.RequestRendering()
				}
			})
		} else if s.flag{
			s.flag = false
			s.componentData.TextColor = s.color

			s.componentData.Redraw()
			engine.RequestRendering()
		}


		return true
	})
}

func (s *Selection) GetName() string {
	return engine.NameOfComponent(s)
}