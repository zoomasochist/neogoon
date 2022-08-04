package effects

import (
	"math/rand"
	"sync/atomic"
	"time"

	"golang.design/x/clipboard"
)

func OverwriteClipboard(annoyanceController <-chan int, decReady *int32) {
	doAnnoyances := false

	atomic.AddInt32(decReady, -1)

	for {
		select {
		case status := <-annoyanceController:
			if status == StartEffects {
				doAnnoyances = true
			} else if status == StopEffects {
				doAnnoyances = false
			}
		default:
			if doAnnoyances {
				// Effect code
				if c.Annoyances.OverwriteClipboard.Chance > rand.Intn(100) {
					text := s.Texts[rand.Intn(len(s.Texts))]
					clipboard.Write(clipboard.FmtText, []byte(text))
				}

				time.Sleep(time.Duration(c.Annoyances.Rate) * time.Second)
			}
		}
	}
}
