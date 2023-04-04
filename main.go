package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	os := runtime.GOOS
	var frameless bool = true

	switch os {
	case "windows":
		frameless = true
	case "darwin":
		frameless = false
	case "linux":
		frameless = true
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options

	err := wails.Run(&options.App{
		Title:     "PDF Strip",
		Width:     1024,
		Height:    768,
		MinWidth:  520,
		MinHeight: 520,
		Frameless: frameless,
		Mac: &mac.Options{
			TitleBar: mac.TitleBarHiddenInset(),
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
