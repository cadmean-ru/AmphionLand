package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"strings"
)

type SearchAlgorithm struct {
	engine.ComponentImpl
	textInput *builtin.NativeInputView
	searchText string
	searchResult map[string] string
	components map[string] string
	counter int
}

func (s *SearchAlgorithm) OnInit(ctx engine.InitContext){
	s.ComponentImpl.OnInit(ctx)

	s.textInput = s.SceneObject.GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.NativeInputView").
		(*builtin.NativeInputView)

	s.textInput.SetHint("Поиск компонентов...")
	s.components = map[string]string{
		"Поле ввода текста" : "github.com/cadmean-ru/amphion/engine/builtin.NativeInputView",
		"Изображение" : "github.com/cadmean-ru/amphion/engine/builtin.ImageView",
	}
	s.searchResult = map[string]string{}

	s.counter = 1

	for i, child := range s.SceneObject.GetChildByName("search results").GetChildren() {
		child.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
			engine.LogDebug("Clicked on position %d", i)
			return true
		}))
	}

	s.textInput.SetOnTextChangeListener(func(text string) {
		s.searchText = s.textInput.GetText()

		for name, component := range s.components {
			if strings.HasPrefix(name, s.searchText) {
				s.searchResult[name] = component
			}
		}

		if s.counter != 1 {
			s.counter = 1
		}

		for name, _ := range s.searchResult {
			textObj := s.SceneObject.GetChildByName("search results").
				GetChildByName(fmt.Sprintf("search result %d", s.counter))
			textView := textObj.
				GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView)
			textView.SetText(name)

			s.counter++
		}

		for i := s.counter; i <= 5; i++ {
			textObj := s.SceneObject.GetChildByName("search results").
				GetChildByName(fmt.Sprintf("search result %d", i))
			textView := textObj.
				GetComponentByName("github.com/cadmean-ru/amphion/engine/builtin.TextView").(*builtin.TextView)
			textView.SetText("")
		}
	})
}

func (s *SearchAlgorithm) GetName() string {
	return engine.NameOfComponent(s)
}