package game

// TODO: Sensitivy
// TODO: Wall check
// TODO: Directional Input

import (
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
	case rl.IsKeyDown(rl.KeyA):
		cx.Player.X -= 0.01
		return
	case rl.IsKeyDown(rl.KeyD):
		cx.Player.X += 0.01
		return
	case rl.IsKeyDown(rl.KeyS):
		cx.Player.Y -= 0.01
		return
	case rl.IsKeyDown(rl.KeyW):
		cx.Player.Y += 0.01
		return
	}
}

func mouseMov(cx *Context) {
	cx.PlayerLookingAt += float64(rl.GetMouseDelta().X * 0.1 * rl.GetFrameTime())
}

// My tears goes here.
func (cx *Context) Process() {
	cx.ProcessRays()
	cx.FindWalls()

	control(cx)
	mov(cx)
	mouseMov(cx)
}
