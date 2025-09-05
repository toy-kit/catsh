package upgrade

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"catsh/global"
	"catsh/types"
	"runtime"
	"time"
)

var (
	GitSource  = []string{"https://github.com", "https://gitee.com"}
	Repository = ""
)

func gitPing(ctx context.Context, done chan struct{}, source string, url string) {
	client := &http.Client{}
	req, _ := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v%v/raw/master/README.md", source, url), nil)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	global.GitSource = source
	close(done)
}

func getRepository(repository string) string {
	uri, err := url.Parse(repository)
	if err != nil {
		return repository
	}
	if global.GitSource != "" {
		return global.GitSource + uri.Path
	}
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	for _, source := range GitSource {
		go gitPing(ctx, done, source, uri.Path)
	}
	timeout := time.After(3 * time.Second)
	select {
	case <-done:
		cancel()
		if global.GitSource != "" {
			return global.GitSource + uri.Path
		} else {
			return repository
		}
	case <-timeout:
		cancel()
		return repository
	}
}

func GetLatest() (types.WailsConfig, error) {
	var wailsConfig types.WailsConfig
	resp, err := http.Get(fmt.Sprintf("%s/raw/master/wails.json", getRepository(global.WailsConfig.Repository)))
	if err != nil {
		return wailsConfig, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return wailsConfig, err
	}
	err = json.Unmarshal(body, &wailsConfig)
	return wailsConfig, err
}

func Download(wailsConfig types.WailsConfig) (string, error) {
	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
	res, err := http.Get(fmt.Sprintf("%v/releases/download/%v/%v%v", getRepository(wailsConfig.Repository), wailsConfig.Info.ProductVersion, wailsConfig.Name, ext))
	if err != nil {
		return "", err
	}
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(filepath.Dir(exePath), fmt.Sprintf("%v%v", wailsConfig.Info.ProductVersion, filepath.Ext(exePath)))
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	_, wErr := io.Copy(file, res.Body)
	if wErr != nil {
		return "", wErr
	}
	defer res.Body.Close()
	defer file.Close()
	return fileName, nil
}

func Install(appPath string) error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	return exec.Command(appPath, "upgrade", exePath).Start()
}

func Upgrade(appPath string) error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	f, err := os.ReadFile(exePath)
	if err != nil {
		return err
	}
	if err := os.WriteFile(appPath, f, 0755); err != nil {
		return err
	}
	go func() {
		time.Sleep(time.Second * 1)
		defer os.Remove(exePath)
		os.Exit(0)
	}()
	return exec.Command(appPath).Start()
}
