package annoyances

import (
	"fmt"
	"math/rand"
	"time"

	"golang.design/x/clipboard"
)

func OverwriteClipboard(annoyanceController <-chan int) {
	doAnnoyances := false

	for {
		select {
		case status := <-annoyanceController:
			if status == StartAnnoyances {
				doAnnoyances = true
				fmt.Println("Going now!!!!!")
			} else if status == StopAnnoyances {
				doAnnoyances = false
			}
		default:
			if doAnnoyances {
				if c.Annoyances.OverwriteClipboard.Chance > rand.Intn(100) {
					text := s.Texts[rand.Intn(len(s.Texts))]
					fmt.Println("Rolled", text)
					clipboard.Write(clipboard.FmtText, []byte(text))
				} else {
					fmt.Println("Roll failed")
				}

				time.Sleep(time.Duration(c.Annoyances.SecondsPerTick) * time.Second)
			}
		}
	}
}
