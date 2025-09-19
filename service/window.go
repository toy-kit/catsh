package service

import (
	"catsh/global"
	"catsh/types"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type WindowService struct{}

// theme background color
var bgColor = map[string][]uint8{
	"light": {255, 255, 255, 1},
	"dark":  {23, 23, 26, 1},
}

func (s *WindowService) NewWindow(options types.WindowOptions) {
	if win, ok := global.WindowManager.GetByName(options.Name); ok {
		win.Restore()
		win.Focus()
		return
	}
	rgba := bgColor[global.Cfg.Theme]
	global.WindowOptions[options.Name] = options
	global.WindowManager.NewWithOptions(application.WebviewWindowOptions{
		Name:                     options.Name,
		Title:                    options.Title,
		Width:                    options.Width,
		Height:                   options.Height,
		MinWidth:                 options.MinWidth,
		MinHeight:                options.MinHeight,
		URL:                      options.URL,
		BackgroundColour:         application.NewRGBA(rgba[0], rgba[1], rgba[2], rgba[3]),
		Frameless:                true,
		ContentProtectionEnabled: true,
		DisableResize:            !options.Resizable,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
	})
}
