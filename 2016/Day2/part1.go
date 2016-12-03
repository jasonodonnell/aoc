package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func read_file(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return instructions
}

func create_matrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	numpad := 1
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = numpad
			numpad += 1
		}
	}
	return matrix
}

func rules(direction string, size int, row int, col int) (int, int) {
	if direction == "L" {
		if (col - 1) > -1 {
			col = col - 1
		}
	} else if direction == "R" {
		if (col + 1) < size {
			col = col + 1
		}
	} else if direction == "U" {
		if (row - 1) > -1 {
			row = row - 1
		}
	} else {
		if (row + 1) < size {
			row = row + 1
		}
	}
	return row, col
}

func main() {
	size := 3
	row := 1
	col := 1

	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	instructions := read_file(*filenamePtr)
	matrix := create_matrix(size)

	for _, element := range instructions {
		for _, direction := range element {
			row, col = rules(string(direction), size, row, col)
		}
		fmt.Printf("%d", matrix[row][col])
	}
	fmt.Println("")
}
