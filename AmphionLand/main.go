package main

import (
	"AmphionLand/components"
	"github.com/cadmean-ru/amphion/engine"
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
}
