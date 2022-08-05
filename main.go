package main

import (
	"fmt"
	config "neogoon/config"
	effects "neogoon/effects"
	set "neogoon/set"
	"os"
	"path/filepath"
	"runtime"

	_ "embed"

	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"
)

//go:embed icon.ico
var icon []byte

func main() {
	if len(os.Args) != 1 && os.Args[1] == "--silent" {
		var c config.Config
		var s set.Set
		prev, err := LoadPreviousSettings()

		ErrIsFatal(err)
		err = config.Load(&c, prev.Config)
		ErrIsFatal(err)
		err = set.Load(&s, prev.Set)
		ErrIsFatal(err)

		effects.Start(&c, &s)
		<-chan int(nil)
	} else {
		systray.Run(Main, func() {})
	}
}

func Main() {
	var c config.Config
	var s set.Set
	var err error

	systray.SetTitle("Neogoon")
	systray.SetIcon(icon)
	systrayStart := systray.AddMenuItem("Start", "Start Neogoon with the configured settings.")
	systray.AddSeparator()

	systrayCurrentConfig := systray.AddMenuItem("Loaded config: (none)", "")
	systrayCurrentPackage := systray.AddMenuItem("Loaded set: (none)", "")
	systrayCurrentConfig.Disable()
	systrayCurrentPackage.Disable()

	prev, err := LoadPreviousSettings()
	ErrIsFatal(err)

	var configPath string
	var setPath string

	if prev.Config == "" {
		systrayStart.Disable()
	} else {
		configPath = prev.Config
		filename := filepath.Base(prev.Config)
		systrayCurrentConfig.SetTitle(fmt.Sprintln("Loaded config:", filename))
	}

	if prev.Set != "" {
		setPath = prev.Set
		filename := filepath.Base(prev.Set)
		systrayCurrentPackage.SetTitle(fmt.Sprintln("Loaded set:", filename))
	}

	systray.AddSeparator()
	systrayLoadConfig := systray.AddMenuItem("Load config", "Load Neogoon Configuration (.toml)")
	systrayLoadPackage := systray.AddMenuItem("Load set", "Load Edgeware Package (.zip)")
	systray.AddSeparator()
	systrayQuit := systray.AddMenuItem("Quit", "Stop Neogoon")

	dialog.Message("Neogoon is ready - load configurations and packages via the tray icon.").Info()

	for {
		select {
		// "Start" was clicked
		case <-systrayStart.ClickedCh:
			systrayStart.SetTitle("Running")
			systrayStart.Disable()

			err = config.Load(&c, configPath)
			ErrIsFatal(err)

			if setPath != "" {
				err = set.Load(&s, setPath)
				ErrIsFatal(err)
			}

			effects.Start(&c, &s)
		// "Load Config" was clicked
		case <-systrayLoadConfig.ClickedCh:
			configPath, err = dialog.File().Filter("Neogoon Configuration File", "toml").Load()
			if err != nil {
				if err == dialog.ErrCancelled {
					continue
				}

				ErrIsFatal(err)
			}

			filename := filepath.Base(configPath)
			systrayCurrentConfig.SetTitle(fmt.Sprintln("Loaded config:", filename))
			systrayStart.Enable()

			err = SaveSettings(configPath, setPath)
			ErrIsFatal(err)

		// "Load Set" was clicked
		case <-systrayLoadPackage.ClickedCh:
			setPath, err = dialog.File().Filter("Neogoon Set", "zip").Load()
			if err != nil {
				if err == dialog.ErrCancelled {
					continue
				}

				ErrIsFatal(err)
			}

			filename := filepath.Base(setPath)
			systrayCurrentPackage.SetTitle(fmt.Sprintln("Loaded set:", filename))

			err = SaveSettings(configPath, setPath)
			ErrIsFatal(err)

		// "Quit" was clicked
		case <-systrayQuit.ClickedCh:
			systray.Quit()
			runtime.Goexit()
		}
	}
}

func ErrIsFatal(err error) {
	if err != nil {
		dialog.Message("From main: ", err.Error()).Error()
		os.Exit(1)
	}
}
