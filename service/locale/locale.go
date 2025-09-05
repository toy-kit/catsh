package locale

import (
	"embed"
	"encoding/json"
	"fmt"
	"path"
	"catsh/global"
	"strings"

	"github.com/tidwall/gjson"
)

var localeJson = ""

func Load(assets embed.FS) error {
	files, err := assets.ReadDir("locales")
	if err != nil {
		return err
	}
	locales := make(map[string]map[string]any)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		b, err := assets.ReadFile(path.Join("locales", file.Name()))
		if err != nil {
			return err
		}
		var data map[string]any
		err = json.Unmarshal(b, &data)
		if err != nil {
			return err
		}
		locales[strings.TrimSuffix(file.Name(), ".json")] = data
	}
	b, err := json.Marshal(locales)
	if err != nil {
		return err
	}
	localeJson = string(b)
	return nil
}

func T(keys string) string {
	return gjson.Get(localeJson, fmt.Sprintf("%v.%v", global.Cfg.Locale, keys)).String()
}

func GetLocales() map[string]map[string]any {
	var locales map[string]map[string]any
	json.Unmarshal([]byte(localeJson), &locales)
	return locales
}
