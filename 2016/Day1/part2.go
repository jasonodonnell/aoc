package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func read_file(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return ""
	}
	return string(data[:])
}

func direction_change(facing int, direction string) int {
	if direction == "R" {
		return (facing + 1) % 4
	} else {
		return (facing + 4 - 1) % 4
	}
}

func walk(x int, y int, facing int, steps int) (int, int) {
	if facing == 0 {
		return x, y + steps
	} else if facing == 1 {
		return x + steps, y
	} else if facing == 2 {
		return x, y - steps
	} else {
		return x - steps, y
	}
}

func main() {
	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	data := read_file(*filenamePtr)
	result := strings.Split(data, ", ")

	x, y := 0, 0

	facing := 0 // 0 North, 1 East, 2 South, 3 West
	visit_twice := make(map[string]int)

	for i := range result {
		directions := strings.Replace(result[i], "\n", "", -1)
		direction := directions[0:1]
		steps, err := strconv.Atoi(directions[1:len(directions)])
		if err != nil {
			fmt.Println(err)
			return
		}

		facing = direction_change(facing, direction)
		for j := 1; j <= steps; j++ {
			x, y = walk(x, y, facing, 1)
			key := fmt.Sprintf("%d,%d", x, y)
			if _, ok := visit_twice[key]; ok {
				fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
				os.Exit(0)
			} else {
				visit_twice[key] = 1
			}
		}
	}
}
