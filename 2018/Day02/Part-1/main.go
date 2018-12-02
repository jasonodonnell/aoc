package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var ids []string

func init() {
	filePath := flag.String("file", "../input.txt", "Path to input file")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Could not open file: %s %s", err, *filePath)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}
}

func main() {
	var doubles int
	var triples int

	for _, id := range ids {
		doubles += checksum(2, id)
		triples += checksum(3, id)
	}

	fmt.Println(doubles * triples)
}

func checksum(matches int, id string) int {
	var count int
	var ids map[string]int
	ids = make(map[string]int)

	for _, letter := range id {
		ids[string(letter)]++
	}

	for _, checksum := range ids {
		if checksum == matches {
			count++
			break
		}
	}

	return count
}
