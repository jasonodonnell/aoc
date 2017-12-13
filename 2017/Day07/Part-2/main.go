package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	t "github.com/jasonodonnell/AdventOfCode/2017/Day07/tower"
)

var tower t.Tower

func init() {
	tower.Programs = make(map[string]*t.Program)
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

		program.Weight, err = strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Fatalf("Error parsing weight: %s", err)
		}

		if len(fields) >= 3 {
			for _, v := range fields[2:len(fields)] {
				program.Children = append(program.Children, v)
			}
			tower.Bases = append(tower.Bases, program.Name)
		}
		tower.Programs[program.Name] = &program
	}
}

func main() {
	fmt.Println(tower.FindUnbalanced(tower.FindBase()))
}

func formatLine(s string) string {
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "->", "", -1)
	return s
}
