package service

import (
	"catsh/global"
	"catsh/internal/config"
	"catsh/internal/locale"
	"catsh/types"
)

type AppDataService struct{}

func (s *AppDataService) Load() types.AppData {
	current := global.WindowManager.Current()
	return types.AppData{
		AppConfig:     global.AppConfig,
		Config:        global.Cfg,
		Locales:       locale.GetLocales(),
		WindowOptions: global.WindowOptions[current.Name()],
	}
}

func (s *AppDataService) SaveConfig(cfg types.Config) {
	config.SaveConf(cfg)
}
