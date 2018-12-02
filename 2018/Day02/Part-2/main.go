package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

var ids []string

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
		ids = append(ids, scanner.Text())
	}
}

func main() {
	sort.Strings(ids)
	for i := 0; i < len(ids)-1; i++ {
		invalid, common := findBox(ids[i], ids[i+1])
		if invalid == 1 {
			fmt.Println(common)
			break
		}
	}
}

func findBox(a string, b string) (int, string) {
	var invalid int
	var common string
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			common += string(a[i])
		}
		if a[i] != b[i] {
			invalid++
		}
	}
	return invalid, common
}
