package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var frequencies []int

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
		frequency, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Could not convert frequency to int: %s", err)
		}
		frequencies = append(frequencies, frequency)
	}
}

func main() {
	var frequency int
	var results map[int]int
	results = make(map[int]int)

	for {
		for _, input := range frequencies {
			frequency += input
			results[frequency]++
			if results[frequency] == 2 {
				fmt.Println(frequency)
				os.Exit(0)
			}
		}
	}
}
