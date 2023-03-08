package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '🏀'

		frames = 1200
		speed  = time.Second / 20
	)

	var (
		px, py int
		vx, vy = 1, 1

		cell rune
	)

	playground := make([][]bool, width)
	for column := range playground {
		playground[column] = make([]bool, height)
	}

	buf := make([]rune, 0, width*height)

	screen.Clear()

	for i := 0; i < frames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := range playground[0] {
			for x := range playground {
				playground[x][y] = false
			}
		}

		playground[px][py] = true

		buf = buf[:0]

		for y := range playground[0] {
			for x := range playground {
				cell = cellEmpty
				if playground[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(speed)
	}
}
