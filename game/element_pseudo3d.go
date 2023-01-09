package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderPseudo3D(cx *Context) {
	winh := float64(rl.GetScreenHeight())

	for p, tile := range cx.Walls {
		go fmt.Printf("Drawing wall at %v which is %d\n", p, tile)
		column_height := winh / p.Delta
		Retangle(int(float64(p.ScreenPosition)), int(winh/2-column_height/2), int(column_height), 1, tile.AsColor())
	}
}
