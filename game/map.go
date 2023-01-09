package game

import (
	"io"
	"os"
)

const (
	Nothing = iota
	FilledBlue
	FilledRed
	FilledGreen
)

func (cx *Context) LoadMap(fileName string) (err error) {
	file, err := os.Open(fileName)

	cx.Map = make([][]int, 16)

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

		case '\n':

			println()
			line++
		}

	}
}
