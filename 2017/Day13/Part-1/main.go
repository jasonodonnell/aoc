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
		line := formatLine(scanner.Text())
		fields := strings.Fields(line)
		k, _ := strconv.Atoi(fields[0])
		v, _ := strconv.Atoi(fields[1])
		l.Depth = v
		firewall.Layers[k] = &l
		if k > largest {
			largest = k
		}
	}
}

func main() {
	var severity int
	for i := 0; i <= largest; i++ {
		if val, ok := firewall.Layers[i]; ok {
			if val.ScannerPos == 0 {
				fmt.Printf("%d Busted\n", i)
				severity += val.Depth * i
			}
		}
		firewall.Move()
	}
	fmt.Println(severity)
}

func formatLine(s string) string {
	s = strings.Replace(s, ":", "", -1)
	return s
}
