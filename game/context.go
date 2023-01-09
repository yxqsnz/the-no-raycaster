package game

type Player struct {
	X, Y float64
}

type Ray struct {
	X, Y float64
}

type Context struct {
	Map        [][]int
	MapHeight  int
	MapWidth   int
	FunnyMode  bool
	MapDeScale int
	Rays       []Ray
	Player     Player
}
