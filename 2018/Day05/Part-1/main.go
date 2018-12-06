package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jasonodonnell/AdventOfCode/2018/Day05/polymer"
)

var polymers []string

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
		polymers = append(polymers, scanner.Text())
	}
}

func main() {
	for _, unit := range polymers {
		p := polymer.New(unit)
		p.React()
		fmt.Println(len(p.Unit))
	}
}
