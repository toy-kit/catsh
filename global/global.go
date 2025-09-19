package global

import (
	"catsh/types"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var (
	WindowManager *application.WindowManager
	WindowOptions = make(map[string]types.WindowOptions)
	AppConfig     types.AppConfig
	Cfg           types.Config
)
