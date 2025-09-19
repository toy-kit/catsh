package upgrade

import (
	"catsh/global"
	"catsh/types"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/stretchr/testify/assert/yaml"
)

var (
	GitSource    = []string{"https://github.com", "https://gitee.com"}
	Source       = ""
	latestConfig = types.AppConfig{}
)

func gitAppConfig(ctx context.Context, done chan struct{}, source string, url string) {
	client := &http.Client{}
	req, _ := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v%v/raw/master/app.yml", source, url), nil)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if body, err := io.ReadAll(resp.Body); err == nil {
		yaml.Unmarshal(body, &latestConfig)
	}
	Source = source
	close(done)
}

func getRepository(repository string) string {
	uri, err := url.Parse(repository)
	if err != nil {
		return repository
	}
	if Source != "" {
		return Source + uri.Path
	}
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	for _, source := range GitSource {
		go gitAppConfig(ctx, done, source, uri.Path)
	}
	timeout := time.After(3 * time.Second)
	select {
	case <-done:
		cancel()
		if Source != "" {
			return Source + uri.Path
		} else {
			return repository
		}
	case <-timeout:
		cancel()
		return repository
	}
}

func Latest() types.AppConfig {
	getRepository(global.AppConfig.Repository.Url)
	return latestConfig
}

func Download(appConfig types.AppConfig) (string, error) {
	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
	res, err := http.Get(fmt.Sprintf("%v/releases/download/%v/%v-%v%v", getRepository(appConfig.Repository.Url), appConfig.Info.Version, appConfig.Info.ProductName, runtime.GOOS, ext))
	if err != nil {
		return "", err
	}
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(filepath.Dir(exePath), fmt.Sprintf("%v%v", appConfig.Info.Version, filepath.Ext(exePath)))
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
