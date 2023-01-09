package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func control(cx *Context) {
	switch {
	case rl.IsKeyPressed(rl.KeyF10):
		cx.FunnyMode = !cx.FunnyMode
	case rl.IsKeyPressed(rl.KeyF6):
		cx.MapDeScale++
	case rl.IsKeyPressed(rl.KeyF5):
		if cx.MapDeScale != 0 {
			cx.MapDeScale--

			if cx.MapDeScale == 0 {
				cx.MapDeScale = 1
			}
		}
	}
}

func mov(cx *Context) {
	switch {
	case rl.IsKeyDown(rl.KeyLeft):
		cx.Player.X += 0.1
	case rl.IsKeyDown(rl.KeyRight):
		cx.Player.X -= 0.1
	case rl.IsKeyDown(rl.KeyUp):
		cx.Player.Y += 0.1
	case rl.IsKeyDown(rl.KeyDown):
		cx.Player.Y -= 0.1
	}
}

// My tears goes here.
func (cx *Context) Process() {
	cx.ProcessRays()
	cx.FindWalls()

	fmt.Printf("Walls: %v\n", cx.Walls)

	control(cx)
	mov(cx)
}
