package effects

import (
	"math/rand"
	"time"

	"github.com/ncruces/zenity"
)

func Notifications(annoyanceController <-chan int) {
	sendNotifications := false

	for {
		select {
		case status := <-annoyanceController:
			if status == StartEffects {
				sendNotifications = true
			} else if status == StopEffects {
				sendNotifications = false
			}
		default:
			if sendNotifications {
				// Effect code
				if c.Annoyances.Notifications.Chance > rand.Intn(100) {
					notification := s.AllTexts[rand.Intn(len(s.AllTexts))]
					zenity.Notify(notification,
						zenity.Title("Neogoon"))
				}

				time.Sleep(time.Duration(c.Annoyances.Rate) * time.Second)
			}
		}
	}
}
