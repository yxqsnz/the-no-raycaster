package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (cx *Context) GetRectSize() (int, int) {
	sh, sw := rl.GetScreenHeight(), rl.GetScreenWidth()
	return sh / cx.MapHeight, sw / cx.MapWidth
}
