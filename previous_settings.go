package main

import (
	"encoding/json"
	config "neogoon/config"
	set "neogoon/set"
	"os"
	"path/filepath"
)

type PreviousSettings struct {
	Config string
	Set    string
}

func LoadPreviousSettings(c *config.Config, s *set.Set) (PreviousSettings, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return PreviousSettings{}, err
	}

	confPath := filepath.Join(dir, "neogoon.json")

	if _, err = os.Stat(confPath); os.IsNotExist(err) {
		return PreviousSettings{}, nil
	} else if err != nil {
		return PreviousSettings{}, err
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

	if err = config.Load(c, previousSettings.Config); err != nil {
		return PreviousSettings{}, err
	}

	if err = set.Load(s, previousSettings.Set); err != nil {
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

	w, err := os.OpenFile(confPath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(b)
	return err
}
