package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"

	"github.com/jasonodonnell/AdventOfCode/2018/Day04/guard"
)

var lines []string

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
		lines = append(lines, scanner.Text())
	}
}

func main() {
	sort.Strings(lines)
	guard.AnalyzeShifts(lines)
}
