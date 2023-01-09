package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (cx *Context) Render() {
	rl.ClearBackground(rl.LightGray)

	RenderRays(cx)
	RenderPseudo3D(cx)

	RenderMap(cx)
	RenderPlayerInMap(cx)
	rl.DrawFPS(20, 20)
}
