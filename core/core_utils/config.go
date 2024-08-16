package core_utils

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func GetAppDir() string {
	appExe, _ := os.Executable()
	appDir := filepath.Dir(appExe)

	return appDir
}

func LoadYamlFile(path string, out any) error {
	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, out)

	if err != nil {
		return err
	}

	return nil
}

func LoadJsonFile(path string, out any) error {
	file, err := os.ReadFile(path)

	if err != nil {
		LogError("LoadJsonFile:%s", err.Error())
		return err
	}

	return json.Unmarshal(file, out)
}

func SaveJsonFile(path string, out any) error {
	bytes, err := json.MarshalIndent(out, "", "  ")

	if err != nil {
		LogError("SaveJsonFile:%s", err.Error())
		return err
	}

	return os.WriteFile(path, bytes, os.ModePerm)
}
