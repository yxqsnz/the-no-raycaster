package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderMap(cx *Context) {
	sh, sw := rl.GetScreenHeight(), rl.GetScreenWidth()
	rect_h, rect_w := sh/cx.MapHeight, sw/cx.MapWidth

	for h, row := range cx.Map {
		for w, col := range row {
			draw := func(col color.RGBA) {
				if cx.FunnyMode {
					go Retangle(w*rect_h/3, h*rect_w/3, 30, 30, col)
				} else {
					Retangle(w*rect_h/3, h*rect_w/3, 30, 30, col)

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
