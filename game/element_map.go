package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderMap(cx *Context) {
	rect_h, rect_w := cx.GetRectSize()

	for h, row := range cx.Map {
		for w, col := range row {
			draw := func(col color.RGBA) {
				render := func() {
					Retangle((w * rect_w / cx.MapDeScale), (h*rect_h)/cx.MapDeScale, rect_h, rect_w, col)
				}

				if cx.FunnyMode {
					go render()
				} else {
					render()
				}
			}

			switch col {
			case Nothing:
				draw(rl.Black)
			case FilledBlue:
				draw(rl.Blue)
			case FilledRed:
				draw(rl.Red)
			case FilledGreen:
				draw(rl.Green)
			}
		}
	}
}
