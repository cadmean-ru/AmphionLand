package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
)

type RadioButton struct {
	engine.ComponentImpl
	length int
	checkedName string
	buttons map[int] string
}

func (s *RadioButton) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
	s.length = 3
	s.checkedName = ""
	s.SceneObject.AddComponent(&builtin.GridLayout{Cols: 2, Rows: s.length})
	s.buttons = make(map[int] string)

	for i := 0; i < s.length; i++ {
		engine.LogDebug("i: %i", i)
		circle := engine.NewSceneObject("circle" + string(rune(i)))
		circle.AddComponent(builtin.NewShapeView(4))
		circle.AddComponent(builtin.NewRectBoundary())
		circle.Transform.Size = a.Vector3{X: 50, Y: 50, Z: 1}
		s.buttons[i] = circle.GetName()
		circle.AddComponent(builtin.NewEventListener(engine.EventMouseDown,func(mouseDownEvent engine.AmphionEvent) bool {
			SetActive(s, mouseDownEvent.Data.(engine.MouseEventData).SceneObject.GetName())
			return true
		}))

		textView := engine.NewSceneObject("textView" + string(rune(i)))
		textView.AddComponent(builtin.NewTextView("i am radio button hello"))
		textView.Transform.Size = a.Vector3{X: 200, Y: 100, Z: 1}

		s.SceneObject.AddChild(circle)
		s.SceneObject.AddChild(textView)
	}
}

func (s *RadioButton) GetName() string {
	return engine.NameOfComponent(s)
}

func SetActive(s *RadioButton, name string) {
	engine.LogDebug("%s", s.checkedName)
	for i := 0; i < len(s.buttons); i++ {
		if s.buttons[i] == s.checkedName {
			engine.LogDebug("found checked")
			checked := engine.NewSceneObject("checked")
			checked.AddComponent(builtin.NewShapeView(4))
			checked.Transform.Size = a.Vector3{X: 25, Y: 25, Z: 2}

			s.SceneObject.FindObjectByName(s.buttons[i]).AddChild(checked)
		} else {
			engine.LogDebug("fuck you")
			s.SceneObject.FindObjectByName(s.buttons[i]).RemoveAllChildren()
		}
	}
}
