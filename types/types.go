package types

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
