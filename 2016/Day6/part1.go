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
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lines
}

type LetterMap struct {
    Key   string
    Value int
}

func main() {
	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	signals := read_file(*filenamePtr)

    const limit = 8
    count := 0

    for count < limit {
        letters := make(map[string]int)
        var repeated_letter LetterMap
        repeated_letter.Value = 0
        for _, each := range signals {
            letters[string(each[count])] += 1
            if letters[string(each[count])] > repeated_letter.Value {
                repeated_letter.Value = letters[string(each[count])]
                repeated_letter.Key = string(each[count])
            }
        }
        fmt.Printf("%s", repeated_letter.Key)
        count++
        repeated_letter.Key = ""
        repeated_letter.Value = 0
    }
    fmt.Println("")
}
