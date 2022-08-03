package annoyances

import (
	"math/rand"
	config "neogoon/config"
	set "neogoon/set"
	"time"
)

const (
	StopAnnoyances = iota
	StartAnnoyances
)

var c *config.Config
var s *set.Set

func Start(cfg *config.Config, p *set.Set) {
	rand.Seed(time.Now().Unix())

	c = cfg
	s = p
	annoyanceController := make(chan int)
	go OverwriteClipboard(annoyanceController)

	annoyanceController <- StartAnnoyances
	// Waits forever
	<-chan int(nil)
}
