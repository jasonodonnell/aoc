package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	m "github.com/jasonodonnell/AdventOfCode/2017/Day5/maze"
)

var maze m.Maze

func init() {
	filePath := flag.String("file", "../input.txt", "Path to input file")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Could not open file: %s %s", err, *filePath)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Error parsing int: %s", err)
		}
		maze.Instructions = append(maze.Instructions, num)
	}
	maze.Position = 0
}

func main() {
	for {
		if maze.Position >= len(maze.Instructions) {
			break
		}
		maze.Move(maze.Instructions[maze.Position], false)
	}
	fmt.Println(maze.Steps)
}
