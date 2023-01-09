package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderPseudo3D(cx *Context) {
	winh := float64(rl.GetScreenHeight())

	for p, tile := range cx.Walls {
		column_height := winh / p.Delta
		render := func() {
			Retangle(int(float64(p.ScreenPosition)), int(winh/2-column_height/2), int(column_height), 1, tile.AsColor())
		}

		if cx.FunnyMode {
			go render()
		} else {
			render()
		}
	}
}
