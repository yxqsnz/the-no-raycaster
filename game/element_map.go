package game

import (
	"image/color"
	"math"
)

func RenderMap(cx *Context) {
	rect_h, rect_w := cx.GetRectSize()

	for h, row := range cx.Map {
		for w, col := range row {
			draw := func(col color.RGBA) {
				render := func() {
					Retangle((w * rect_w / cx.MapDeScale), (h*rect_h)/cx.MapDeScale, int(math.Ceil(float64(rect_w)/float64(cx.MapDeScale))), int(math.Ceil(float64(rect_w)/float64(cx.MapDeScale))), col)
				}

				if cx.FunnyMode {
					go render()
				} else {
					render()
				}
			}

			draw(col.AsColor())
		}
	}
}
