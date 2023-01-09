package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (cx *Context) Render() {
	rl.ClearBackground(rl.LightGray)

	RenderMap(cx)
	RenderPlayerInMap(cx)
	RenderRays(cx)
	RenderPseudo3D(cx)

	rl.DrawFPS(20, 20)
}
