package service

import (
	"catsh/internal/locale"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/icons"
)

func NewService() []application.Service {
	return []application.Service{
		application.NewService(&AppDataService{}),
		application.NewService(&UpgradeService{}),
	}
}

func Alert(message string) {
	dialog := application.InfoDialog()
	dialog.SetTitle(locale.T("dialog.title"))
	dialog.SetIcon(icons.ApplicationDarkMode256)
	dialog.SetMessage(message)
	dialog.Show()
}

func QuestionDialog(message string) bool {
	dialog := application.QuestionDialog()
	dialog.SetTitle(locale.T("dialog.title"))
	dialog.SetMessage(message)
	dialog.AddButton("Save").OnClick(func() {
		// Handle save
	})
	saveButton := dialog.AddButton("Don't Save")
	dialog.SetDefaultButton(saveButton)
	dialog.Show()
	return true
}
