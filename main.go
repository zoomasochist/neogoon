package main

import (
	annoyances "neogoon/annoyances"
	config "neogoon/config"
	set "neogoon/set"
	"os"

	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"
)

func main() {
	systray.Run(Main, func() {})
}

func Main() {
	dialog.Message("Neogoon is ready - load configurations and packages via the tray icon.").Info()

	systray.SetTitle("Neogoon")
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
	var configLoaded bool
	var packageLoaded bool

	for {
		select {
		case <-systrayStart.ClickedCh:
			annoyances.Start(&c, &s)
		case <-systrayLoadConfig.ClickedCh:
			fullPath, err := LoadConfig(&c)
			if err != nil {
				dialog.Message(err.Error()).Error()
				continue
			}
			configLoaded = true
			filename := filepath.Base(fullPath)
			systrayCurrentConfig.SetTitle(fmt.Sprintln("Loaded config:", filename))
			if configLoaded && packageLoaded {
				systrayStart.Enable()
			}

		case <-systrayLoadPackage.ClickedCh:
			fullPath, err := LoadPackage(&s)
			if err != nil {
				dialog.Message(err.Error()).Error()
				continue
			}
			packageLoaded = true
			filename := filepath.Base(fullPath)
			systrayCurrentPackage.SetTitle(fmt.Sprintln("Loaded package:", filename))
			if configLoaded && packageLoaded {
				systrayStart.Enable()
			}

		case <-systrayQuit.ClickedCh:
			systray.Quit()
			os.Exit(0)
		}
	}

}

func LoadConfig(c *config.Config) (string, error) {
	filename, err := dialog.File().Filter("Neogoon Configuration File", "toml").Load()
	if err != nil {
		return "", err
	}

	_, err = toml.DecodeFile(filename, &c)
	if err != nil {
		return "", err
	}

	fmt.Printf("%+v", c)

	return filename, err
}

func LoadPackage(s *set.Set) (string, error) {
	s.Texts = []string{"Puppy", "Doggy", "Smalley"}

	return "", nil
}
