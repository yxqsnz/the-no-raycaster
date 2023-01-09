package game

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (cx *Context) ProcessRays() {
	fov := math.Pi / 3.
	cx.Rays = make([]Position, rl.GetScreenWidth())
	cx.Colisions = make([]Colision, rl.GetScreenWidth())

	for i := 0; i < rl.GetScreenWidth(); i++ {
		angle := cx.PlayerLookingAt - fov/2. + fov*float64(i)/float64(rl.GetScreenWidth())

		for distance := 0.; distance < 20.; distance += .5 {
			rayX := cx.Player.X + distance*math.Cos(angle)
			rayY := cx.Player.Y + distance*math.Sin(angle)

			if cx.Map[int(rayY)][int(rayX)] != Nothing {
				cx.Colisions = append(cx.Colisions, Colision{Position{rayX, rayY}, i, distance})
				break
			}

			cx.Rays = append(cx.Rays, Position{rayX, rayY})
		}
	}
}
