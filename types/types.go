package types

import (
	"fmt"
	"strings"
)

type WailsConfigAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type WailsConfigInfo struct {
	CompanyName    string `json:"companyName"`
	ProductName    string `json:"productName"`
	ProductVersion string `json:"productVersion"`
	Copyright      string `json:"copyright"`
	Comments       string `json:"comments"`
}

// WailsConfig wails.json
type WailsConfig struct {
	Name           string            `json:"name"`
	OutputFilename string            `json:"outputfilename"`
	Repository     string            `json:"repository"`
	HomePage       string            `json:"homepage"`
	Author         WailsConfigAuthor `json:"author"`
	Info           WailsConfigInfo   `json:"info"`
}

// Config config.ini
type Config struct {
	Theme  string `json:"theme" ini:"theme"`
	Locale string `json:"locale" ini:"locale"`
}

type LoadData struct {
	WailsConfig WailsConfig               `json:"wails_config"`
	Config      Config                    `json:"config"`
	Locales     map[string]map[string]any `json:"locales"`
}

type Database struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Protocol  string            `json:"protocol"`
	Transport string            `json:"transport"`
	User      string            `json:"user"`
	Pass      string            `json:"pass"`
	Host      string            `json:"host"`
	DbName    string            `json:"dbname"`
	Options   map[string]string `json:"options"`
}

func (d Database) URL() string {
	url := fmt.Sprintf("%v+%v://%v:%v@%v/%v?", d.Protocol, d.Transport, d.User, d.Pass, d.Host, d.DbName)
	options := []string{}
	for k, v := range d.Options {
		options = append(options, fmt.Sprintf("%v=%v", k, v))
	}
	return fmt.Sprintf("%v?%v", url, strings.Join(options, "&"))
}
