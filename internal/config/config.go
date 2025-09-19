package config

import (
	"catsh/global"
	"catsh/types"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/ini.v1"
)

// 获取配置文件
func getConfFile() string {
	filename := filepath.Join("."+global.AppConfig.Info.CompanyName, global.AppConfig.Info.ProductName, "conf.ini")
	dir := ""
	if u, err := user.Current(); err == nil {
		dir = u.HomeDir
	}
	confFile := filepath.Join(dir, filename)
	os.MkdirAll(filepath.Dir(confFile), os.ModePerm)
	return confFile
}
func LoadConf() error {
	cfg := types.Config{
		Theme:  "light",
		Locale: "zh",
	}
	i, err := ini.LoadSources(ini.LoadOptions{SkipUnrecognizableLines: true}, getConfFile())
	if err != nil {
		i = ini.Empty()
	}
	err = i.MapTo(&cfg)
	global.Cfg = cfg
	return err
}

func SaveConf(cfg types.Config) {
	global.Cfg = cfg
	i := ini.Empty()
	err := ini.ReflectFrom(i, &cfg)
	if err != nil {
		return
	}
	i.SaveTo(getConfFile())
}
