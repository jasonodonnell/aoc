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
Loop:
	for _, claim := range claims {
		for i := 0; i < claim.Tall; i++ {
			for j := 0; j < claim.Wide; j++ {
				if len(fabric.Matrix[claim.X+j][claim.Y+i]) > 1 {
					continue Loop
				}
			}
		}
		fmt.Println(claim.ID)
	}
}
