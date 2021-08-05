package main

import (
	"AmphionLand/components"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/rpc"
)

type AppDelegate struct {
	engine.AppDelegateImpl
}

func (d *AppDelegate) OnAppLoaded() {
	cm := engine.GetComponentsManager()
	cm.RegisterEventHandler(components.OnClick)

	engine.LogDebug("App loaded")

	rpc.Initialize("http://localhost:4200")
}

func (d *AppDelegate) OnAppStopping() {
	engine.LogDebug("App stopping")
}