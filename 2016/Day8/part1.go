package main

import (
    "bufio"
	"flag"
	"fmt"
	"os"
    "strconv"
    "strings"
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

func create_matrix(row int, col int) [][]int {
	matrix := make([][]int, col)
	for i := range matrix {
		matrix[i] = make([]int, row)
	}
	return matrix
}

func rect_rules(wide int, tall int, matrix [][]int) [][]int {
    for i := 0; i < tall; i++ {
        for j := 0; j < wide; j++ {
            matrix[i][j] = 1
        }
    }
    return matrix
}

func rotate_column(col int, rotate int, max_col int, matrix [][]int) [][]int {
    rotated := make([]int, max_col)
    for i := 0; i < max_col; i++ {
        index := (i + rotate) % max_col
        rotated[index] = matrix[i][col]
    }
    
    for j := 0; j < max_col; j++ {
        matrix[j][col] = rotated[j]
    }
    return matrix
}

func rotate_row(row int, rotate int, max_row int, matrix [][]int) [][]int {
    rotated := make([]int, max_row)
    for i := 0; i < max_row; i++ {
        index := (i + rotate) % max_row
        rotated[index] = matrix[row][i]
    }

    for j := 0; j < max_row; j++ {
        matrix[row][j] = rotated[j]
    }
    return matrix
}

func main() {
	row := 50 
	col := 6
    sum := 0

	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	file_input := read_file(*filenamePtr)
	matrix := create_matrix(row, col)

    for _, each := range file_input {
        instructions := strings.Fields(each)
        if instructions[0] == "rect" {
            rect := strings.Split(instructions[1], string('x'))
            wide, _ := strconv.Atoi(rect[0])
            tall, _ := strconv.Atoi(rect[1])
            matrix = rect_rules(wide, tall, matrix)
        } else {
            rotate_count, _ := strconv.Atoi(strings.Split(instructions[2], string('='))[1])
            rotate, _ := strconv.Atoi(instructions[4])
            if instructions[1] == "column" {
                matrix = rotate_column(rotate_count, rotate, col, matrix)
            } else {
                matrix = rotate_row(rotate_count, rotate, row, matrix)
            }
        }
    }

    for _, matrix_row := range matrix {
        for _, matrix_col := range matrix_row {
            if matrix_col == 1 {
                sum++
            }

        }
    }
	fmt.Println(sum)
}
