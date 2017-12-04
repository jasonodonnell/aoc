package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jasonodonnell/AdventOfCode/2017/Day4/runesort"
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
	for _, passphrase := range passphrases {
		for k, word := range passphrase {
			if k+1 != len(passphrase) {
				if anagram(word, passphrase[k+1:]) == true {
					invalid++
					break
				}
			}
		}
	}
	fmt.Println(len(passphrases) - invalid)
}

func anagram(word string, compare []string) bool {
	word = runesort.Sort(strings.ToLower(word))
	for _, s := range compare {
		if len(word) != len(s) {
			continue
		}
		s = runesort.Sort(strings.ToLower(s))
		if word == s {
			return true
		}
	}
	return false
}
