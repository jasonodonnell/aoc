package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jasonodonnell/AdventOfCode/2017/Day03/spiral"
)

var board *spiral.Spiral
var currentPosition *spiral.Point
var data *int

func init() {
	data = flag.Int("data", 1, "Data to move")
	flag.Parse()

	board = &spiral.Spiral{
		Board: map[int]spiral.Point{
			1: spiral.Point{
				X: 0.0,
				Y: 0.0,
			},
		},
		Sums: map[spiral.Point]int{
			spiral.Point{X: 0.0, Y: 0.0}: 1,
		},
		Position:  1, // Start Position
		Direction: 0, // Direction to Move
		Repeat:    1, // Repeat direction state (grows by 1 every two direction change)
	}
	currentPosition = &spiral.Point{X: 0.0, Y: 0.0}
}

func main() {
	for board.Position <= *data {
		if board.Position == 1 {
			board.Position++
			continue
		}
		for i := 0; i < 2; i++ {
			for j := 0; j < board.Repeat; j++ {
				currentPosition = board.Move(*currentPosition)
				board.Sums[*currentPosition] = currentPosition.SumAdjacent(board)
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
