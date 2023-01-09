package game

type Position struct {
	X, Y float64
}

type ContretePosition struct {
	X, Y int
}

type Context struct {
	Map        [][]int
	Walls      map[ContretePosition]int
	Colisions  []Position
	MapHeight  int
	MapWidth   int
	FunnyMode  bool
	MapDeScale int
	Rays       []Position
	Player     Position
}
