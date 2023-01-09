package game

func (cx *Context) FindWalls() {
	cx.Walls = make(map[ContretePosition]int)

	for _, colision := range cx.Colisions {
		x, y := int(colision.X), int(colision.Y)
		p := ContretePosition{
			x, y,
		}

		cx.Walls[p] = cx.Map[y][x]
	}
}
