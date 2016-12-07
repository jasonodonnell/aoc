package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func read_file(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lines
}

func get_int(room string) int {
	re := regexp.MustCompile("[0-9]+")
	sector_id, _ := strconv.Atoi(re.FindAllString(room, -1)[0])
	return sector_id
}

func get_string(room string) string {
	re := regexp.MustCompile("[a-z]+")
	result := re.FindAllString(room, -1)
	var str bytes.Buffer

	for _, each := range result {
		str.WriteString(each)
	}

	return str.String()
}

func main() {
	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	rooms := read_file(*filenamePtr)

    letters := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e": 5,
        "f": 6,
        "g": 7,
        "h": 8,
        "i": 9,
        "j": 10,
        "k": 11,
        "l": 12,
        "m": 13,
        "n": 14,
        "o": 15,
        "p": 16,
        "q": 17,
        "r": 18,
        "s": 19,
        "t": 20,
        "u": 21,
        "v": 22,
        "w": 23,
        "x": 24,
        "y": 25,
        "z": 26,
    }

	for _, room := range rooms {
		hash := strings.Split(room, "[")

		sector_id := get_int(hash[0])

		result := strings.Replace(string(hash[0]), "-", " ", -1)
		var buffer bytes.Buffer
		for _, each := range result {
			if string(each) == " " {
				buffer.WriteString(" ")
				continue
			} else if _, err := strconv.Atoi(string(each)); err == nil {
				break
			} else {
				increment := (letters[string(each)] + sector_id) % 26
                for k,v := range letters {
                    if v == increment {
				        buffer.WriteString(string(k))
                        break
                    }
                }
			}
		}
		fmt.Println(buffer.String(), sector_id)
	}
}
