package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	m "github.com/jasonodonnell/AdventOfCode/2017/Day6/memory"
)

var memory m.Memory

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
		nums := strings.Fields(scanner.Text())
		for _, v := range nums {
			num, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			memory.Banks = append(memory.Banks, num)
		}
	}
	memory.Blocks = make(map[string]int)
}

func main() {
	var repeated string
	for {
		if memory.Stringify(memory.Banks) == repeated {
			break
		}
		if memory.BankExists() {
			if repeated == "" {
				repeated = memory.Stringify(memory.Banks)
				memory.RedistributionCount = 0
				memory.Blocks = make(map[string]int)
			}
		}
		memory.BlockRedistribution()
	}
	fmt.Println(memory.RedistributionCount)
}
