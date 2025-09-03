package main

import (
	"context"
	"catsh/global"
	"catsh/internal/config"
	"catsh/internal/locale"
	"catsh/internal/upgrade"
	"catsh/types"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Load
func (a *App) Load() types.LoadData {
	// Load locale
	if err := locale.Load(locales); err != nil {
		println("Error:", err.Error())
	}
	return types.LoadData{
		WailsConfig: global.WailsConfig,
		Config:      global.Cfg,
		Locales:     locale.GetLocales(),
	}
}

// GetConfig
func (a *App) SaveConf(cfg types.Config) {
	config.SaveConf(cfg)
}

// CheckUpgrade
func (a *App) CheckUpgrade() {
	wailsConfig, err := upgrade.GetLatest()
	if err != nil {
		a.Alert(err.Error())
		return
	}
	if wailsConfig.Info.ProductVersion == global.WailsConfig.Info.ProductVersion {
		a.Alert(locale.T("upgrade.noUpdate"))
		return
	}
	if ok := a.QuestionDialog(locale.T("upgrade.update")); !ok {
		return
	}
	appPath, err := upgrade.Download(wailsConfig)
	if err != nil {
		a.Alert(err.Error())
		return
	}
	if ok := a.QuestionDialog(locale.T("upgrade.downloaded")); !ok {
		return
	}
	if err := upgrade.Install(appPath); err != nil {
		a.Alert(err.Error())
		return
	} else {
		runtime.Quit(a.ctx)
	}
}

// OpenAbout
func (a *App) OpenAbout() {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   locale.T("titlebar.about"),
		Message: locale.T("description"),
	})
}
func (a *App) QuestionDialog(message string) bool {
	result, _ := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   locale.T("dialog.title"),
		Message: message,
	})
	return result == "Yes"
}
func (a *App) Alert(message string) {
	runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   locale.T("dialog.title"),
		Message: message,
	})
}
