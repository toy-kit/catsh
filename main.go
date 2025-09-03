package main

import (
	"embed"
	"encoding/json"
	"os"
	"catsh/global"
	"catsh/internal/config"
	"catsh/internal/upgrade"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed wails.json
var wailsConfig []byte

//go:embed all:locales
var locales embed.FS

var bgColor = map[string][]uint8{
	"light": {255, 255, 255, 1},
	"dark":  {23, 23, 26, 1},
}

func main() {
	// Read wails.json
	if err := json.Unmarshal(wailsConfig, &global.WailsConfig); err != nil {
		println("Error:", err.Error())
	}
	// Load config
	if err := config.LoadConf(); err != nil {
		println("Error:", err.Error())
	}

	// Upgrade
	if len(os.Args) == 3 && os.Args[1] == "upgrade" {
		if err := upgrade.Upgrade(os.Args[2]); err != nil {
			println("Error:", err.Error())
		}
		return
	}

	// get background color
	rgba := bgColor[global.Cfg.Theme]

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            global.WailsConfig.Name,
		Width:            900,
		Height:           650,
		MinWidth:         800,
		MinHeight:        600,
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: rgba[0], G: rgba[1], B: rgba[2], A: rgba[3]},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			DisablePinchZoom: true,
		},
		Mac: &mac.Options{
			DisableZoom: true,
		},
		Linux: &linux.Options{},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
