package main

import (
	config "neogoon/config"
	effects "neogoon/effects"
	set "neogoon/set"
	"os"

	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"
)

//go:embed icon.ico
var icon []byte

func main() {
	if len(os.Args) != 1 && os.Args[1] == "--silent" {
		var c config.Config
		var s set.Set
		_, err := LoadPreviousSettings(&c, &s)
		if err != nil {
			dialog.Message(err.Error()).Error()
		}

		effects.Start(&c, &s)
		<-chan int(nil)
	} else {
		systray.Run(Main, func() {})
	}
}

func Main() {
	dialog.Message("Neogoon is ready - load configurations and packages via the tray icon.").Info()

	systray.SetTitle("Neogoon")
	systray.SetIcon(icon)
	systrayStart := systray.AddMenuItem("Start", "Start Neogoon with the loaded settings.")
	systrayStart.Disable()
	systrayCurrentConfig := systray.AddMenuItem("Loaded config: (none)", "")
	systrayCurrentPackage := systray.AddMenuItem("Loaded package: (none)", "")
	systrayCurrentConfig.Disable()
	systrayCurrentPackage.Disable()
	systrayLoadConfig := systray.AddMenuItem("Load Config", "Load Neogoon Configuration (.toml)")
	systrayLoadPackage := systray.AddMenuItem("Load Package", "Load Edgeware Package (.zip)")
	systray.AddSeparator()
	systrayQuit := systray.AddMenuItem("Quit", "Stop Neogoon")

	var c config.Config
	var s set.Set
	var configPath string
	var setPath string
	var running bool
	var configLoaded bool

	previous, err := LoadPreviousSettings(&c, &s)
	if err != nil {
		dialog.Message(err.Error()).Error()
	}

	if len(previous.Set) != 0 {
		filename := filepath.Base(previous.Set)
		systrayCurrentPackage.SetTitle(fmt.Sprintln("Loaded set:", filename))
	}
	if len(previous.Config) != 0 {
		filename := filepath.Base(previous.Config)
		configLoaded = true
		systrayStart.Enable()
		systrayCurrentConfig.SetTitle(fmt.Sprintln("Loaded config:", filename))
	}

	for {
		select {
		case <-systrayStart.ClickedCh:
			running = true
			systrayStart.SetTitle("Running")
			systrayStart.Disable()
			effects.Start(&c, &s)
		case <-systrayLoadConfig.ClickedCh:
			configPath, err = LoadConfig(&c)
			filename := filepath.Base(configPath)
			if err != nil {
				dialog.Message(err.Error()).Error()
			} else {
				systrayCurrentConfig.SetTitle(fmt.Sprintln("Loaded config:", filename))
				if configLoaded && !running {
					systrayStart.Enable()
				}
			}

		case <-systrayLoadPackage.ClickedCh:
			setPath, err = LoadSet(&s)
			if err != nil {
				dialog.Message(err.Error()).Error()
			}
			filename := filepath.Base(setPath)
			systrayCurrentPackage.SetTitle(fmt.Sprintln("Loaded set:", filename))

		case <-systrayQuit.ClickedCh:
			err = SaveSettings(configPath, setPath)
			if err != nil {
				dialog.Message(err.Error()).Error()
			}
			systray.Quit()
			os.Exit(0)
		}
	}
}

func LoadConfig(c *config.Config) (string, error) {
	configPath, err := dialog.File().Filter("Neogoon Configuration File", "toml").Load()
	if err != nil {
		return "", err
	}

	err = config.Load(c, configPath)
	if err != nil {
		return "", err
	}

	// Normalise paths
	c.DriveFiller.Root = filepath.FromSlash(c.DriveFiller.Root)

	return configPath, nil
}

func LoadSet(s *set.Set) (string, error) {
	setPath, err := dialog.File().Filter("Neogoon Set", "zip").Load()
	if err != nil {
		return setPath, err
	}

	err = set.Load(s, setPath)
	if err != nil {
		return setPath, err
	}

	return setPath, nil
}
