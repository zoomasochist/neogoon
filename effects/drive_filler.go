package effects

import (
	"fmt"
	"math/rand"
	booru "neogoon/effects/booru"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DriveFiller(annoyanceController <-chan int) {
	fillDrive := false

	imageNames := c.DriveFiller.Filenames
	b, err := booru.FromString(c.Downloader.Booru)
	if err != nil {
		Fault(err.Error())
	}

	directories, err := EnumeratePaths(c.DriveFiller.Root)
	if err != nil {
		Fault(err.Error())
	}

	for {
		select {
		case status := <-annoyanceController:
			if status == StartEffects {
				fillDrive = true
			} else if status == StopEffects {
				fillDrive = false
			}
		default:
			if fillDrive {
				// Effect code
				image := b.Next()
				path := directories[rand.Intn(len(directories))]
				randomName := fmt.Sprintf("%s%d.%s",
					imageNames[rand.Intn(len(imageNames))],
					time.Now().Unix(),
					image.Ext)
				writeTo := filepath.Join(path, randomName)

				err := os.WriteFile(writeTo, image.Bytes, 0644)
				if err != nil {
					Fault(err.Error())
				}

				time.Sleep(time.Duration(c.DriveFiller.Rate) * time.Millisecond)
			}
		}
	}
}

// This is kinda slow atm
func EnumeratePaths(path string) ([]string, error) {
	var r []string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return r, err
	}

	path = strings.Replace(path, "~", homeDir, 1)

	r = append(r, path)

	return r, filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if os.IsPermission(err) || !info.IsDir() {
				return nil
			}

			if err != nil {
				return err
			}

			r = append(r, path)
			return nil
		})
}
