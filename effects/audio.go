package effects

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
)

func Audio(annoyanceController <-chan int) {
	audio := false
	currentlyPlaying := false

	ctrl := &beep.Ctrl{}
	var streamer beep.StreamSeekCloser
	var format beep.Format

	for {
		select {
		case status := <-annoyanceController:
			if status == StartEffects {
				audio = true
				ctrl.Paused = false
			} else if status == StopEffects {
				audio = false
				ctrl.Paused = true
			}
		default:
			if audio {
				// Effect code
				if !currentlyPlaying && c.Annoyances.Audio.Chance > rand.Intn(100) {
					audioFilePath := s.Audio[rand.Intn(len(s.Audio))]
					f, err := os.Open(audioFilePath)
					if err != nil {
						panic(err)
					}

					components := strings.Split(audioFilePath, ".")
					switch components[len(components)-1] {
					case "mp3":
						streamer, format, _ = mp3.Decode(f)
					case "wav":
						streamer, format, _ = wav.Decode(f)
					case "flac":
						streamer, format, _ = flac.Decode(f)
					case "ogg":
						streamer, format, _ = vorbis.Decode(f)
					}

					speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
					ctrl.Streamer = streamer
					ctrl.Paused = false
					volume := &effects.Volume{
						Streamer: ctrl,
						Base:     2,
						// Thanks Gray
						Volume: (float64((c.Annoyances.Audio.Volume * 2)) - 200) / 100,
					}
					speaker.Play(beep.Seq(volume, beep.Callback(func() {
						time.Sleep(time.Duration(c.Annoyances.Rate) * time.Second)
						currentlyPlaying = false
					})))

					if c.Annoyances.Audio.MaxPlaytime != 0 {
						go func() {
							time.Sleep(time.Duration(c.Annoyances.Audio.MaxPlaytime) * time.Second)
							currentlyPlaying = false
							ctrl.Paused = true
						}()
					}

					currentlyPlaying = true
				}
			}
		}
	}
}
