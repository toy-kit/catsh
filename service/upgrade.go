package service

import (
	"catsh/global"
	"catsh/internal/locale"
	"catsh/internal/upgrade"
)

type UpgradeService struct{}

func (s *UpgradeService) CheckAndInstall() {
	latestConfig := upgrade.Latest()
	if latestConfig.Info.Version == global.AppConfig.Info.Version {
		Alert(locale.T("upgrade.noUpdate"))
		return
	}
	if ok := QuestionDialog(locale.T("upgrade.update")); !ok {
		return
	}
	appPath, err := upgrade.Download(latestConfig)
	if err != nil {
		Alert(err.Error())
		return
	}
	if ok := QuestionDialog(locale.T("upgrade.downloaded")); !ok {
		return
	}
	if err := upgrade.Install(appPath); err != nil {
		Alert(err.Error())
		return
	}
}
