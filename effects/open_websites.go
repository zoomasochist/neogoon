package effects

import (
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"time"
)

func OpenWebsites(annoyanceController <-chan int) {
	openWebsites := false

	for {
		select {
		case status := <-annoyanceController:
			if status == StartEffects {
				openWebsites = true
			} else if status == StopEffects {
				openWebsites = false
			}
		default:
			if openWebsites {
				// Effect code
				if c.Annoyances.OpenWebsites.Chance > rand.Intn(100) {
					var err error
					url := s.Urls[rand.Intn(len(s.Urls))]

					// I don't like shelling out like this because I can't know what the runtime
					// system has installed. Windows will always have rundll32 and macOS will
					// always have open, but Linux won't always have xdg-open. A solution that isn't
					// this would be nice.
					switch runtime.GOOS {
					case "linux":
						err = exec.Command("xdg-open", url).Start()
					case "windows":
						err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
					case "darwin":
						err = exec.Command("open", url).Start()
					default:
						err = fmt.Errorf("unsupported platform for opening URLs")
					}
					if err != nil {
						panic(err)
					}
				}

				time.Sleep(time.Duration(c.Annoyances.Rate) * time.Second)
			}
		}
	}
}
