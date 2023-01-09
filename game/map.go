package game

import (
	"image/color"
	"io"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile int

const (
	Nothing Tile = iota
	FilledBlue
	FilledRed
	FilledGreen
)

func (t Tile) AsColor() color.RGBA {
	switch t {
	case Nothing:
		return rl.Black
	case FilledBlue:
		return rl.Blue
	case FilledRed:
		return rl.Red
	case FilledGreen:
		return rl.Green
	default:
		return color.RGBA{255, 255, 255, 255}
	}
}

func (cx *Context) LoadMap(fileName string) (err error) {
	file, err := os.Open(fileName)

	cx.Map = make([][]Tile, 16)

	if err != nil {
		return err
	}

	defer file.Close()

	buf := make([]byte, 1)
	line := 0

	for {

		_, err := file.Read(buf)

		if err == io.EOF {

			return nil
		}

		if err != nil {
			return err
		}

		switch buf[0] {
		case ' ':
			print(" ")
			cx.Map[line] = append(cx.Map[line], Nothing)
		case 'b':
			print("b")
			cx.Map[line] = append(cx.Map[line], FilledBlue)
		case 'r':
			print("r")
			cx.Map[line] = append(cx.Map[line], FilledRed)
		case 'g':
			print("g")
			cx.Map[line] = append(cx.Map[line], FilledGreen)
		case '*':
			print("*")
			cx.Map[line] = append(cx.Map[line], FilledBlue)
		case '\n':
			println()
			line++
		}
	}

}
