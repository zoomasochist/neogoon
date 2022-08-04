package effects

import (
	"math/rand"
	config "neogoon/config"
	set "neogoon/set"
	"os"
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

	var waitingOn int32
	c = cfg
	s = p
	annoyanceController := make(chan int)

	if c.Annoyances.OverwriteClipboard.Chance > 0 {
		waitingOn++
		go OverwriteClipboard(annoyanceController, &waitingOn)
	}
	if c.DriveFiller.Enabled {
		waitingOn++
		go DriveFiller(annoyanceController, &waitingOn)
	}

	// I don't think this is the best solution but it functions well so eh
	for waitingOn != 0 {
		time.Sleep(200 * time.Millisecond)
	}

	annoyanceController <- StartEffects
}

func Fault(message string) {
	dialog.Message(message).Error()
	os.Exit(1)
}
