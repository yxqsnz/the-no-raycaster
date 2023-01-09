package game

import rl "github.com/gen2brain/raylib-go/raylib"

func RenderRays(cx *Context) {
	rect_h, rect_w := cx.GetRectSize()

	for _, ray := range cx.Rays {
		render := func() {
			rl.DrawRectangle(int32(ray.X*float64(rect_w)/float64(cx.MapDeScale)), int32(ray.Y*float64(rect_h)/float64(cx.MapDeScale)), 1, 1, rl.Pink)
		}

		if cx.FunnyMode {
			go render()
		} else {
			render()
		}
	}
}
