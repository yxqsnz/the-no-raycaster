package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (cx *Context) ensureMap() {
	for _, line := range cx.Map {
		if len(line) == 0 {
			break
		}

		if len(line) < cx.MapWidth {
			panic("Line is smaller than others!")
		} else if len(line) > cx.MapWidth {
			cx.MapWidth = len(line)
		}

		cx.MapHeight++
	}
}

func (cx *Context) Prepare() {
	err := cx.LoadMap("map.txt")

	if err != nil {
		panic(err)
	}

	cx.FunnyMode = false
	cx.Player.X = 3.456
	cx.Player.Y = 2.345
	cx.PlayerLookingAt = 1.523
	cx.MapDeScale = 6

	cx.ensureMap()
	fmt.Printf("\nMap (%dx%d): %v\n", cx.MapHeight, cx.MapWidth, cx.Map[:cx.MapHeight])
	rl.DisableCursor()
}
