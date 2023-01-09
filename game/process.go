package game

import rl "github.com/gen2brain/raylib-go/raylib"

// My tears goes here.
func (cx *Context) Process() {
	if rl.IsKeyPressed(rl.KeyF10) {
		cx.FunnyMode = !cx.FunnyMode
	}
}
