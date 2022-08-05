package effects

import (
	"math/rand"
	config "neogoon/config"
	set "neogoon/set"
	"runtime"
	"time"

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

	if c.Annoyances.OverwriteClipboard.Chance > 0 {
		go OverwriteClipboard(annoyanceController)
		annoyanceController <- StartEffects
	}

	if c.DriveFiller.Enabled {
		go DriveFiller(annoyanceController)
		annoyanceController <- StartEffects
	}

	if c.Wallpaper.Enabled {
		go SetWallpaper(annoyanceController)
		annoyanceController <- StartEffects
	}
}

func Fault(message string) {
	dialog.Message(message).Error()
	runtime.Goexit()
}
