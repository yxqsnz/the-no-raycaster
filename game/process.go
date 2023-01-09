package game

import rl "github.com/gen2brain/raylib-go/raylib"

// My tears goes here.
func (cx *Context) Process() {
	if rl.IsKeyPressed(rl.KeyF10) {
		cx.FunnyMode = !cx.FunnyMode
	}

	if rl.IsKeyPressed(rl.KeyF5) {
		if cx.MapDeScale != 0 {
			cx.MapDeScale--

			if cx.MapDeScale == 0 {
				cx.MapDeScale = 1
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyF6) {
		cx.MapDeScale++
	}
}
