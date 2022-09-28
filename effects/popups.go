package effects

import (
	"math/rand"
	"neogoon/gui"
	"time"
)

func Popups(annoyanceController <-chan int) {
	doPopups := false

	ui, _ := gui.InitUI()
	for {
		select {
		case status := <-annoyanceController:
			if status == StartEffects {
				doPopups = true
			} else if status == StopEffects {
				doPopups = false
			}
		default:
			if doPopups {
				// Effect code
				if c.Annoyances.Popups.Chance > rand.Intn(100) {
					imagePath := s.Images[rand.Intn(len(s.Images))]
					go ui.Popup(imagePath)
				}

				time.Sleep(time.Duration(c.Annoyances.Rate) * time.Second)
			}
		}
	}
}
