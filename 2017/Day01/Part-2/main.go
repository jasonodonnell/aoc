package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

var dat []byte
var err error

func init() {
	filePath := flag.String("file", "../input.txt", "Path to input file")
	flag.Parse()

	dat, err = ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("Could not read file: %s %s", err, *filePath)
	}
}

func main() {
	var captcha []int
	var sum int

	for _, v := range dat {
		value, err := strconv.Atoi(string(v))
		if err != nil {
			continue
		}
		captcha = append(captcha, value)
	}

	stepSize := len(captcha) / 2

	for i := range captcha {
		j := (i + stepSize) % len(captcha)
		if captcha[i] == captcha[j] {
			sum += captcha[i]
		}
	}
	fmt.Println(sum)
}
