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
	var lines []string
	for scanner.Scan() {
		lines = append(lines , scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lines 
}

func valid_triangle(a int, b int, c int) bool {
    if ( a >= b + c ) || ( b >= a + c ) || ( c >= a + b ) {
        return false
    } else {
        return true
    }
}

func main() {
	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	triangles := read_file(*filenamePtr)
    col_a := []int{}
    col_b := []int{}
    col_c := []int{}
    valid_triangles := 0

	for _, element := range triangles {
        result := strings.Fields(element)

        if result == nil {
            continue
        }
        
        a, _ := strconv.Atoi(result[0])
        b, _ := strconv.Atoi(result[1])
        c, _ := strconv.Atoi(result[2])

        col_a = append(col_a, a)
        col_b = append(col_b, b)
        col_c = append(col_c, c)
    }
  
    col_d := make([]int, len(col_a) + len(col_b) + len(col_c))
    copy(col_d[:], col_a[:])
    copy(col_d[len(col_a):], col_b[:])
    copy(col_d[len(col_a) + len(col_b):], col_c[:])

    for i := 0; i < len(col_d); i=i+3 {
        if valid_triangle(col_d[i], col_d[i+1], col_d[i+2]) {
            valid_triangles += 1
        }
	}
    fmt.Println(valid_triangles)
}
