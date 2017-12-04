package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jasonodonnell/AdventOfCode/2017/Day3/spiral"
)

func main() {
	data := flag.Int("data", 1, "Data to move")
	flag.Parse()

	board := spiral.Spiral{
		Board: map[int]spiral.Point{
			1: spiral.Point{
				X: 0.0,
				Y: 0.0,
			},
		},
		Sums: map[spiral.Point]int{
			spiral.Point{X: 0.0, Y: 0.0}: 1,
		},
		Position:  1,
		Direction: 0,
		Repeat:    1,
	}

	currentPosition := &spiral.Point{X: 0.0, Y: 0.0}

	for board.Position <= *data {
		if board.Position == 1 {
			board.Position++
			continue
		}
		for j := 0; j < 2; j++ {
			for k := 0; k < board.Repeat; k++ {
				currentPosition = board.Move(*currentPosition)
				board.Sums[*currentPosition] = currentPosition.SumAdjacent(&board)
				if board.Sums[*currentPosition] > *data {
					fmt.Println(board.Sums[*currentPosition])
					os.Exit(0)
				}
				board.Position++
			}
			board.Direction = (board.Direction + 1) % 4
		}
		board.Repeat++
	}
}
