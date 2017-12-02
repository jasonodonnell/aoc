package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jasonodonnell/AdventOfCode/2017/Day2/checksum"
)

var checksums []*checksum.Checksum

func init() {
	filePath := flag.String("file", "../input.txt", "Path to input file")
	flag.Parse()

	dat, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Could not open file: %s %s", err, *filePath)
	}
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		checksum := &checksum.Checksum{}
		nums := strings.Fields(scanner.Text())
		for _, v := range nums {
			num, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			checksum.Nums = append(checksum.Nums, num)
		}
		checksums = append(checksums, checksum)
	}
}

func main() {
	checksum := 0
	for _, c := range checksums {
		checksum += c.Difference()
	}
	fmt.Println(checksum)
}
