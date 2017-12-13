package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	s "github.com/jasonodonnell/AdventOfCode/2017/Day09/stream"
)

var stream string

func init() {
	filePath := flag.String("file", "./input.txt", "Path to input file")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Could not open file: %s %s", err, *filePath)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		stream = scanner.Text()
	}
}

func main() {
	s := s.Stream{}
	s.ProcessStream(stream)
	fmt.Println(s.GroupCount, s.GarbageCount)
}
