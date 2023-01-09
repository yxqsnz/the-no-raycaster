package game

import "fmt"

func (cx *Context) ensureMap() {

	for _, line := range cx.Map {
		if len(line) == 0 {
			break
		}

		if len(line) < cx.MapWidth {
			panic("Line is minor than others!")
		} else if len(line) > cx.MapWidth {
			cx.MapWidth = len(line)
		}

		cx.MapHeight++
	}
}

func (cx *Context) Prepare() {
	err := cx.LoadMap("map.txt")

	if err != nil {
		panic(err)
	}

	cx.ensureMap()
	fmt.Printf("\nMap (%dx%d): %v\n", cx.MapHeight, cx.MapWidth, cx.Map[:cx.MapHeight])

}
