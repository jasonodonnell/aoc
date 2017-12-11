package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jasonodonnell/AdventOfCode/2017/Day11/hex"
)

var directions []string

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
		directions = strings.Split(scanner.Text(), ",")
	}
}

func main() {
	var h hex.Hex
	for _, direction := range directions {
		h.Move(direction)
	}
	fmt.Printf("Distance: %.0f, Furthest: %.0f\n", h.Distance, h.Furthest)
}
