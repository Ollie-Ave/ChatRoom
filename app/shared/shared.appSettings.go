package shared

import (
	"encoding/json"
	"fmt"
	"os"
)

var appSettings *AppSettings

type AppSettings struct {
	ApplicationPort     int
	TemplatePath        string
	StaticFileDirectory string
}

func GetAppSettings() (*AppSettings, error) {
	if appSettings == nil {
		var err error

		appSettings, err = loadAppSettings(os.Getenv("APP_ENVIRONMENT"))

		if err != nil {
			return nil, err
		}
	}

	return appSettings, nil
}

func loadAppSettings(environment string) (*AppSettings, error) {
	var appSettings AppSettings

	fileNameToLoad := fmt.Sprintf("./appsettings.%s.json", environment)

	fileData, err := os.ReadFile(fileNameToLoad)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileData, &appSettings)

	if err != nil {
		return nil, err
	}

	return &appSettings, nil
}
