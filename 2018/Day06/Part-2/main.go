package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jasonodonnell/AdventOfCode/2018/Day06/grid"
)

var coordinates []string

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
		coordinates = append(coordinates, scanner.Text())
	}
}

func main() {
	var points []grid.Point
	for _, coordinate := range coordinates {
		var point grid.Point
		n, err := fmt.Sscanf(coordinate, "%d, %d", &point.X, &point.Y)
		if n != 2 {
			log.Fatalf("Could not scan coordinates: %s", err)
		}
		points = append(points, point)
	}

	g := grid.New(points)

	limit := 10000
	fmt.Println(g.RegionSum(limit))
}
