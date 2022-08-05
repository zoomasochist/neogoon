package effects

import (
	"fmt"
	"io"
	"math/rand"
	config "neogoon/config"
	set "neogoon/set"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/emersion/go-autostart"
	"github.com/sqweek/dialog"
)

const (
	StopEffects = iota
	StartEffects
)

var c *config.Config
var s *set.Set

func Start(cfg *config.Config, p *set.Set) {
	rand.Seed(time.Now().Unix())

	c = cfg
	s = p
	annoyanceController := make(chan int)

	path, err := os.UserConfigDir()
	if err != nil {
		Fault(err.Error())
	}

	fp := filepath.Join(path, "gw.exe")

	as := &autostart.App{
		Name:        "Neogoon",
		DisplayName: "Neogoon",
		Exec:        []string{fmt.Sprintf("\"%s\"", fp), "--silent"},
	}

	if c.Infection.StartOnBoot {
		if err := CopyFile(os.Args[0], fp); err != nil {
			Fault(err.Error())
		}
		if !as.IsEnabled() {
			if err := as.Enable(); err != nil {
				Fault(err.Error())
			}
		}
	} else {
		if as.IsEnabled() {
			if err := as.Disable(); err != nil {
				Fault(err.Error())
			}
		}
	}

	if c.Annoyances.OverwriteClipboard.Chance > 0 && len(s.Texts) > 0 {
		go OverwriteClipboard(annoyanceController)
		annoyanceController <- StartEffects
	}

	if c.DriveFiller.Enabled {
		go DriveFiller(annoyanceController)
		annoyanceController <- StartEffects
	}

	if c.Wallpaper.Enabled && len(s.Images) > 0 {
		go SetWallpaper(annoyanceController)
		annoyanceController <- StartEffects
	}
}

func Fault(message string) {
	dialog.Message(message).Error()
	runtime.Goexit()
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}
