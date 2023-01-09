package game

type Player struct {
	X float64
	Y float64
}

type Context struct {
	Map        [][]int
	MapHeight  int
	MapWidth   int
	FunnyMode  bool
	MapDeScale int
	Player     Player
}
