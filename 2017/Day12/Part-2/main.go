package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	p "github.com/jasonodonnell/AdventOfCode/2017/Day12/program"
)

var programs p.Programs

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

	programs.Pipes = make(map[string][]string)
	programs.Visited = make(map[string]bool)
	for scanner.Scan() {
		line := formatLine(scanner.Text())
		fields := strings.Fields(line)
		programs.Pipes[fields[0]] = fields[:len(fields)]
	}
}

func main() {
	groups := 0
	for k := range programs.Pipes {
		if _, ok := programs.Visited[k]; !ok {
			programs.WalkGroup(k)
			groups++
		}
	}
	fmt.Println(groups)
}

func formatLine(s string) string {
	s = strings.Replace(s, "<->", "", -1)
	s = strings.Replace(s, ",", "", -1)
	return s
}
