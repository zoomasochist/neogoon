package effects

import (
	"math/rand"
	"time"

	"golang.design/x/clipboard"
)

func OverwriteClipboard(annoyanceController <-chan int) {
	doAnnoyances := false

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
					text := s.AllTexts[rand.Intn(len(s.AllTexts))]
					clipboard.Write(clipboard.FmtText, []byte(text))
				}

				time.Sleep(time.Duration(c.Annoyances.Rate) * time.Second)
			}
		}
	}
}
