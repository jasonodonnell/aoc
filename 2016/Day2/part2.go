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

func create_matrix(size int) [][]string {
	matrix := [][]string{}
	row1 := []string{"X", "X", "1", "X", "X"}
	row2 := []string{"X", "2", "3", "4", "X"}
	row3 := []string{"5", "6", "7", "8", "9"}
	row4 := []string{"X", "A", "B", "C", "X"}
	row5 := []string{"X", "X", "D", "X", "X"}
	matrix = append(matrix, row1)
	matrix = append(matrix, row2)
	matrix = append(matrix, row3)
	matrix = append(matrix, row4)
	matrix = append(matrix, row5)
	return matrix
}

func rules(direction string, size int, row int, col int, matrix [][]string) (int, int) {
	if direction == "L" {
		if (col-1) > -1 && matrix[row][col-1] != "X" {
			col = col - 1
		}
	} else if direction == "R" {
		if (col+1) < size && matrix[row][col+1] != "X" {
			col = col + 1
		}
	} else if direction == "U" {
		if (row-1) > -1 && matrix[row-1][col] != "X" {
			row = row - 1
		}
	} else {
		if (row+1) < size && matrix[row+1][col] != "X" {
			row = row + 1
		}
	}
	return row, col
}

func main() {
	size := 5
	row := 2
	col := 0

	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	instructions := read_file(*filenamePtr)
	matrix := create_matrix(size)

	for _, element := range instructions {
		for _, direction := range element {
			row, col = rules(string(direction), size, row, col, matrix)
		}
		fmt.Printf("%s", string(matrix[row][col]))
	}
	fmt.Println("")
}
