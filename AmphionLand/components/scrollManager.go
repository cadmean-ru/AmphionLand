//+build js

package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/engine"
	"syscall/js"
)

type Scrolling struct {
	engine.ComponentImpl
}

func (s *Scrolling) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)

	js.Global().Get("document").Call("addEventListener", "wheel", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var event = args[0]
		event.Call("preventDefault")

		var deltaY = event.Get("deltaY").Float()

		engine.LogDebug(fmt.Sprintf("deltaY: %f", deltaY))

		s.SceneObject.Transform.Position.Y += float32(deltaY)

		engine.RequestRendering()

		return nil
	}))
}

func (s *Scrolling) OnUpdate(ctx engine.UpdateContext) {
}

func (s *Scrolling) GetName() string {
	return engine.NameOfComponent(s)
}

