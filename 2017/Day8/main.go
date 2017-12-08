package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jasonodonnell/AdventOfCode/2017/Day8/cpu"
)

var instructions []*cpu.Instruction

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
		field := strings.Fields(scanner.Text())
		instruction, err := cpu.NewInstruction(field)
		if err != nil {
			log.Fatalf("Error parsing instruction: %s", err)
		}
		instructions = append(instructions, instruction)
	}
}

func main() {
	for _, v := range instructions {
		v.ProcessInstruction()
	}
	fmt.Println(cpu.LargestRegister())
	fmt.Println(cpu.Highmark())
}
