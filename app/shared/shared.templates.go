package shared

import "fmt"

func GetTemplatePath(templateName string) string {
	appSettings, err := GetAppSettings()

	if err != nil {
		panic(err)
	}

	templatePath := fmt.Sprintf("%s/%s", appSettings.TemplatePath, templateName)

	return templatePath
}
