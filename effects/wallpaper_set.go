package effects

import (
	"image"
	"math/rand"
	"os"
	"time"

	"github.com/reujab/wallpaper"
	"github.com/sqweek/dialog"
)

var oldWallpaper string

func SetWallpaper(annoyanceController <-chan int) {
	setWallpaper := false

	background, err := wallpaper.Get()
	if err != nil {
		dialog.Message("Couldn't get your old wallpaper - it won't be reset: ", err.Error()).Info()
	}
	oldWallpaper = background

	wallpaperSetMode := AsSetMode(c.Wallpaper.Mode)
	wallpaper.SetMode(wallpaperSetMode)

	var wallpaperPool []string

	if c.Wallpaper.PreferFit {
		for _, img := range s.Images {
			handle, err := os.Open(img)
			if err != nil {
				Fault(err.Error())
			}

			meta, _, err := image.DecodeConfig(handle)
			// image can't decode everything apparently. Whatever.
			if err == image.ErrFormat {
				continue
			}
			if err != nil {
				Fault(err.Error())
			}

			ratio := float32(meta.Width) / float32(meta.Height)
			if ratio > 1.7 && ratio < 2. {
				wallpaperPool = append(wallpaperPool, img)
			}
		}
	}

	if len(wallpaperPool) == 0 {
		wallpaperPool = s.Images
	}

	defer func() { wallpaper.SetFromFile(oldWallpaper) }()

	for {
		select {
		case status := <-annoyanceController:
			if status == StartEffects {
				setWallpaper = true
			} else {
				setWallpaper = false
				wallpaper.SetFromFile(oldWallpaper)
			}
		default:
			if setWallpaper {
				wp := wallpaperPool[rand.Intn(len(wallpaperPool))]
				err := wallpaper.SetFromFile(wp)
				if err != nil {
					Fault(err.Error())
				}

				time.Sleep(time.Duration(c.Wallpaper.Rate) * time.Millisecond)
			}
		}
	}
}

func AsSetMode(s string) wallpaper.Mode {
	switch s {
	case "center":
		return wallpaper.Center
	case "crop":
		return wallpaper.Crop
	case "fit":
		return wallpaper.Fit
	case "span":
		return wallpaper.Span
	case "stretch":
		return wallpaper.Stretch
	case "tile":
		return wallpaper.Tile
	default:
		return wallpaper.Center
	}
}
