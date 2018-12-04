package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jasonodonnell/AdventOfCode/2018/Day03/fabric"
)

var claims []fabric.Claim

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
		claims = append(claims, *fabric.NewClaim(scanner.Text()))
	}
}

func main() {
	fabric := fabric.NewFabric(claims)
	wastedFabric := 0
	for i := 0; i < len(fabric.Matrix); i++ {
		for j := 0; j < len(fabric.Matrix); j++ {
			if len(fabric.Matrix[i][j]) > 1 {
				wastedFabric++
			}
		}
	}
	fmt.Println(wastedFabric)
}
