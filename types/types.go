package types

import (
	"fmt"
	"strings"
)

type AppConfigRepository struct {
	Author   string `json:"author"`
	Email    string `json:"email"`
	Url      string `json:"url"`
	Homepage string `json:"homepage"`
}

type AppConfigInfo struct {
	CompanyName string `json:"companyName"`
	ProductName string `json:"productName"`
	Description string `json:"description"`
	Copyright   string `json:"copyright"`
	Version     string `json:"version"`
}

// AppConfig app.yml
type AppConfig struct {
	Info       AppConfigInfo       `json:"info"`
	Repository AppConfigRepository `json:"repository"`
	OS         string              `json:"os"`
}

// Config config.ini
type Config struct {
	Theme  string `json:"theme" ini:"theme"`
	Locale string `json:"locale" ini:"locale"`
}

type AppData struct {
	AppConfig     AppConfig                 `json:"app_config"`
	Config        Config                    `json:"config"`
	WindowOptions WindowOptions             `json:"window_options"`
	Locales       map[string]map[string]any `json:"locales"`
}

type WindowOptions struct {
	Name      string `json:"name"`
	Title     string `json:"title"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	MinWidth  int    `json:"min_width"`
	MinHeight int    `json:"min_height"`
	Resizable bool   `json:"resizable"`
	URL       string `json:"url"`
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
