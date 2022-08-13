package effects

import (
	"math/rand"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
)

var num = map[string]int{
	"1": 2, "2": 3, "3": 4, "4": 5, "5": 6,
	"6": 7, "7": 8, "8": 9, "9": 10, "0": 11,
}

var alpha = map[string]int{
	"q": 16, "w": 17, "e": 18, "r": 19, "t": 20,
	"y": 21, "u": 22, "i": 23, "o": 24, "p": 25,
	"a": 30, "s": 31, "d": 32, "f": 33, "g": 34,
	"h": 35, "j": 36, "k": 37, "l": 38, "z": 44,
	"x": 45, "c": 46, "v": 47, "b": 48, "n": 49,
	"m": 50,
}

var symbol = map[string]int{
	"-": 12, "_": -12, "=": 13, "+": -13, "[": 26,
	"{": -26, "]": 27, "}": -27, "'": 40, "\"": -40,
	"`": 41, "~": -41, "\\": 43, "|": -43, ",": 51,
	"<": -51, ".": 52, ">": -52, "/": 53, "?": -53,
	" ": 57, "!": -2, "@": -3, "#": -4, "$": -5,
	"%": -6, "^": -7, "&": -8, "*": -9, "(": -10,
	")": -11, ";": 39, ":": -39,
}

func AutoType(annoyanceController <-chan int) {
	autoType := false

	for {
		select {
		case status := <-annoyanceController:
			if status == StartEffects {
				autoType = true
			} else if status == StopEffects {
				autoType = false
			}
		default:
			if autoType {
				// Effect code
				if c.Annoyances.AutoType.Chance > rand.Intn(100) {
					text := s.AllTexts[rand.Intn(len(s.AllTexts))]

					err := TypeKeys(text, c.Annoyances.AutoType.PressEnter)
					if err != nil {
						panic(err)
					}
				}

				time.Sleep(time.Duration(c.Annoyances.Rate) * time.Second)
			}
		}
	}
}

func TypeKeys(s string, pressEnter bool) error {
	kbWrap, err := keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}

	for _, section := range strings.Split(s, " ") {
		chars := strings.Split(section, "")
		for _, c := range chars {
			d, dok := num[c]
			a, aok := alpha[c]
			sym, symok := symbol[c]
			ca, caok := alpha[strings.ToLower(c)]

			switch {
			case aok:
				kbWrap.SetKeys(a)
				kbWrap.Launching()
			case dok:
				kbWrap.SetKeys(d)
				kbWrap.Launching()
			case caok:
				kbWrap.HasSHIFT(true)
				kbWrap.SetKeys(ca)
				kbWrap.Launching()
				kbWrap.HasSHIFT(false)
			case symok:
				kbWrap.SetKeys(sym)
				kbWrap.Launching()
			}
		}
		// Space
		kbWrap.SetKeys(57)
		kbWrap.Launching()
	}

	if pressEnter {
		kbWrap.SetKeys(keybd_event.VK_ENTER)
		kbWrap.Launching()
	}

	return nil
}

/*
func TypeKeys(s string) error {
	kbWrap, err := keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}

	for _, section := range strings.Split(s, " ") {
		chars := strings.Split(section, "")
		for _, c := range chars {
			d, dok := num[c]
			a, aok := alpha[c]
			sym, symok := symbol[c]
			ca, caok := alpha[strings.ToLower(c)]

			switch {
			case aok:
				kbWrap.HasSHIFT(false)
				kbWrap.SetKeys(a)
			case dok:
				kbWrap.HasSHIFT(false)
				kbWrap.SetKeys(d)
			case caok:
				kbWrap.HasSHIFT(true)
				kbWrap.SetKeys(ca)
			case symok:
				kbWrap.HasSHIFT(false)
				kbWrap.SetKeys(sym)
			}
			err = kbWrap.Launching()
			if err != nil {
				return err
			}
		}
		// Space
		kbWrap.SetKeys(57)
		kbWrap.Launching()
	}

	return nil
}
*/
