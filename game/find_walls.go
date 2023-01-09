package game

func (cx *Context) FindWalls() {
	cx.Walls = make(map[Colision]Tile)

	for _, colision := range cx.Colisions {
		x, y := int(colision.At.X), int(colision.At.Y)

		cx.Walls[colision] = cx.Map[y][x]
	}
}
