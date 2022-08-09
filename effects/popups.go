package effects

import (
	"neogoon/virtualscreen"

	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/image/draw"
)

func Popups(annoyanceController <-chan int) {
	doPopups := false

	displayX, displayY := virtualscreen.Resolution()

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
					imgPath := s.Images[rand.Intn(len(s.Images))]
					f, err := os.Open(imgPath)
					if err != nil {
						panic(err)
					}

					m, _, err := image.Decode(f)
					if err != nil {
						panic(err)
					}

					f.Close()

					//topLeft := rand.Intn(displayX)
					//height := rand.Intn(displayY)

					imageX := m.Bounds().Dx()
					imageY := m.Bounds().Dy()

					modX := imageX / displayX
					modY := imageY / displayY

					var maxMod int
					if modX > modY {
						maxMod = modX + 1
					} else {
						maxMod = modY + 1
					}

					imgWidth := (imageX / maxMod) / 2
					imgHeight := (imageY / maxMod) / 2

					resizedImage := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
					// Other options in order of performance/quality ratio:
					// ApproxBiLinear, BiLinear, CatmullRom
					draw.NearestNeighbor.Scale(resizedImage, resizedImage.Rect, m, m.Bounds(),
						draw.Over, nil)

					closeAfter := -1
					if c.Annoyances.Popups.Timeout != 0 {
						closeAfter = c.Annoyances.Popups.Timeout
					}

					go SpawnPopup(resizedImage, imgWidth, imgHeight, closeAfter)
				}

				time.Sleep(time.Duration(c.Annoyances.Rate) * time.Second)
			}
		}
	}
}

func SpawnPopup(img image.Image, width, height, closeAfter int) {
	w := app.NewWindow()
	w.Option(app.Decorated(false))
	w.Option(app.Size(unit.Dp(width), unit.Dp(height)+35))
	w.Option(app.Title("Submit <3"))
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	var closeButtonWidget widget.Clickable
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceEnd,
			}.Layout(gtx,
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						imageOp := paint.NewImageOp(img)
						imageOp.Add(&ops)
						op.Affine(f32.Affine2D{}.Scale(f32.Pt(0, 0), f32.Pt(4, 4)))
						paint.PaintOp{}.Add(&ops)

						return layout.Dimensions{Size: imageOp.Size()}
					},
				),

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						closeButton := material.Button(th, &closeButtonWidget, "Submit <3")
						closeButton.Background = color.NRGBA{0, 0, 0, 255}
						closeButton.CornerRadius = unit.Dp(0)
						return closeButton.Layout(gtx)
					},
				),
			)

			if closeButtonWidget.Clicked() {
				w.Perform(system.ActionClose)
			}

			e.Frame(gtx.Ops)
		}
	}
}
