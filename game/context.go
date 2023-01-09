package game

type Position struct {
	X, Y float64
}

type ContretePosition struct {
	X, Y int
}

type Colision struct {
	At             Position
	ScreenPosition int
	Distance       float64
}

type Context struct {
	Map        [][]Tile
	Walls      map[Colision]Tile
	Colisions  []Colision
	MapHeight  int
	MapWidth   int
	FunnyMode  bool
	MapDeScale int
	Rays       []Position
	Player     Position
}
