package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	t "github.com/jasonodonnell/AdventOfCode/2017/Day7/tower"
)

var tower t.Tower

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
		var program t.Program
		line := formatLine(scanner.Text())
		fields := strings.Fields(line)
		program.Name = fields[0]

		program.Weight, err = strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("Error parsing weight: %s", err)
		}

		if len(fields) >= 3 {
			for _, v := range fields[2:len(fields)] {
				program.Supports = append(program.Supports, v)
			}
			tower.Bases = append(tower.Bases, program.Name)
		}
		tower.Programs = append(tower.Programs, &program)
	}
}

func main() {
	fmt.Println(tower.FindBase())
}

func formatLine(s string) string {
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "->", "", -1)
	return s
}
