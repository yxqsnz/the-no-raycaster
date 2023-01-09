package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Why this module exists? Remove useless convertions from RayLib.

func Retangle(x int, y int, height int, width int, color color.RGBA) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), color)
}
