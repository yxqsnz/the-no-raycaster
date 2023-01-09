package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (cx *Context) Render() {
	rl.ClearBackground(rl.LightGray)

	RenderMap(cx)
	RenderPlayerInMap(cx)
	RenderRays(cx)

	rl.DrawFPS(20, 20)
}
