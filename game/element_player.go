package game

import rl "github.com/gen2brain/raylib-go/raylib"

func RenderPlayerInMap(cx *Context) {
	rect_h, rect_w := cx.GetRectSize()
	Retangle(int(cx.Player.X*float64(rect_w))/cx.MapDeScale, int(cx.Player.Y*float64(rect_h))/cx.MapDeScale, 5, 5, rl.White)
}
