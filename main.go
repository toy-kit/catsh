package main

import (
	"catsh/global"
	"catsh/internal/config"
	"catsh/internal/locale"
	"catsh/internal/upgrade"
	"catsh/service"
	"catsh/types"
	"embed"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
	"gopkg.in/yaml.v3"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

//go:embed app.yml
var appConfig []byte

//go:embed all:locales
var locales embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	// Read app.yml
	if err := yaml.Unmarshal(appConfig, &global.AppConfig); err != nil {
		println("Error:", err.Error())
	}
	global.AppConfig.OS = runtime.GOOS

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

	// Load locale
	if err := locale.Load(locales); err != nil {
		println("Error:", err.Error())
	}

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "Catsh",
		Description: "Catsh",
		Services:    service.NewService(),
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		SingleInstance: &application.SingleInstanceOptions{
			UniqueID: "com.lufuhu.catsh",
			OnSecondInstanceLaunch: func(data application.SecondInstanceData) {
				// Focus the window if needed
				if mainWindow, ok := global.WindowManager.GetByName("main"); ok {
					mainWindow.Restore()
					mainWindow.Focus()
				}
			},
		},
	})

	global.WindowManager = app.Window

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	windowService := service.WindowService{}
	windowService.NewWindow(types.WindowOptions{
		Name:      "main",
		Title:     "Catsh",
		Width:     900,
		Height:    650,
		MinWidth:  800,
		MinHeight: 600,
		URL:       "/",
		Resizable: true,
	})

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Event.Emit("time", now)
			time.Sleep(time.Second * 3)
		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
