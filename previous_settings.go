package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type PreviousSettings struct {
	Config string
	Set    string
}

func LoadPreviousSettings() (PreviousSettings, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return PreviousSettings{}, err
	}

	confPath := filepath.Join(dir, "neogoon.json")

	if !FileExists(confPath) {
		return PreviousSettings{}, nil
	}

	conf, err := os.ReadFile(confPath)
	if err != nil {
		return PreviousSettings{}, err
	}

	var previousSettings PreviousSettings
	err = json.Unmarshal(conf, &previousSettings)
	if err != nil {
		return PreviousSettings{}, err
	}

	return previousSettings, nil
}

func SaveSettings(configPath, setPath string) error {
	b, err := json.Marshal(PreviousSettings{Config: configPath, Set: setPath})
	if err != nil {
		return err
	}

	dir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	confPath := filepath.Join(dir, "neogoon.json")

	w, err := os.Create(confPath)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(b)
	return err
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil || os.IsNotExist(err) {
		return false
	}

	return true
}
