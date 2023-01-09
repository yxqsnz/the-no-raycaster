package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (cx *Context) GetRectSize() (int, int) {
	sh, sw := rl.GetScreenHeight(), rl.GetScreenWidth()
	return sh / cx.MapHeight, sw / cx.MapWidth
}

func Retangle(x int, y int, height int, width int, color color.RGBA) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), color)
}
