package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jasonodonnell/AdventOfCode/2017/Day10/knot"
)

var lengths []byte

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
		lengths = []byte(scanner.Text())
	}
	additional := []byte{17, 31, 73, 47, 23}
	lengths = append(lengths, additional...)
}

func main() {
	k := knot.NewKnot(256)
	for i := 0; i < 64; i++ {
		for _, v := range lengths {
			k.Reverse(int(v))
		}
	}
	fmt.Println(k.Hash())
}
