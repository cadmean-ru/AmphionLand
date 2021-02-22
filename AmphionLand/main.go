package main

import (
	"AmphionLand/components"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"github.com/cadmean-ru/amphion/engine/rpc"
)

func main() {
	runApp()
}

func registerComponents(cm *engine.ComponentsManager) {
	cm.RegisterComponentType(&components.TestComponent{})
	cm.RegisterComponentType(&components.Scrolling{})
	cm.RegisterComponentType(&components.InputField{})
	cm.RegisterComponentType(&components.MainSceneController{})
	cm.RegisterComponentType(&components.Selection{})
	cm.RegisterComponentType(&components.Zooming{})
	cm.RegisterComponentType(&components.LoginSceneController{})
	cm.RegisterComponentType(&builtin.NativeInputView{})
	cm.RegisterComponentType(&components.WidgetGrid{})
	cm.RegisterComponentType(&components.WodgetController{})

	rpc.Initialize("http://localhost:4200")
}
