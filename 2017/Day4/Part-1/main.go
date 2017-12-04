package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var passphrases [][]string

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
		passphrase := strings.Fields(scanner.Text())
		passphrases = append(passphrases, passphrase)
	}
}

func main() {
	invalid := 0
	for _, list := range passphrases {
		passphrase := make(map[string]int)
		for _, phrase := range list {
			if _, ok := passphrase[phrase]; !ok {
				passphrase[phrase]++
			} else {
				invalid++
				break
			}
		}
	}
	fmt.Println(len(passphrases) - invalid)
}
