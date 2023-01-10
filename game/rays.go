package game

import (
	"math"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ty Garicas :pepe_heart:
func (cx *Context) ProcessRays() {
	fov := math.Pi / 3.

	var wg sync.WaitGroup

	for i := 0; i < rl.GetScreenWidth(); i++ {
		wg.Add(1)
		go func(i int) {
			angle := cx.PlayerLookingAt - fov/2. + fov*float64(i)/float64(rl.GetScreenWidth())

			for distance := 0.; distance < 20.; distance += .05 {
				rayX := cx.Player.X + distance*math.Cos(angle)
				rayY := cx.Player.Y + distance*math.Sin(angle)

				if cx.Map[int(rayY)][int(rayX)] != Nothing {
					cx.Colisions[i] = Colision{Position{rayX, rayY}, i, distance}
					break
				}

				cx.Rays[i] = Position{rayX, rayY}
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
