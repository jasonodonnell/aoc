package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jasonodonnell/AdventOfCode/2017/Day10/knot"
)

var lengths []int

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
		field := strings.Split(scanner.Text(), ",")
		for _, v := range field {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			lengths = append(lengths, num)
		}
	}
}

func main() {
	k := knot.NewKnot(256)
	for _, v := range lengths {
		k.Reverse(v)
	}
	fmt.Println(k.List[0] * k.List[1])
}
