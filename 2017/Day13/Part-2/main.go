package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	fire "github.com/jasonodonnell/AdventOfCode/2017/Day13/firewall"
)

var firewall fire.Firewall
var largest int

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

	firewall.Layers = make(map[int]*fire.Layer)
	for scanner.Scan() {
		var l fire.Layer
		line := strings.Split(scanner.Text(), ": ")
		k, _ := strconv.Atoi(line[0])
		v, _ := strconv.Atoi(line[1])
		l.Depth = v
		firewall.Layers[k] = &l
		if k > largest {
			largest = k
		}
	}
}

func main() {
	delay := 1
Loop:
	for {
		for i := 0; i <= largest; i++ {
			if val, ok := firewall.Layers[i]; ok {
				// Double the depth, trim off the ends, check
				// mod against current time.  Zero means busted
				if (delay+i)%((val.Depth*2)-2) == 0 {
					delay++
					continue Loop
				}
			}
		}
		break
	}
	fmt.Println(delay)
}
