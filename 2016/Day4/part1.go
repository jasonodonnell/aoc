package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
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

func sort_by_letter_count(letters map[string]int) PairList {
	pl := make(PairList, len(letters))
	i := 0
	for k, v := range letters {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	rooms := read_file(*filenamePtr)
	total_sum := 0

	for _, room := range rooms {
		letters := make(map[string]int)
		hash := strings.Split(room, "[")

		encrypted_name := get_string(hash[0])
		sector_id := get_int(hash[0])
		checksum := get_string(hash[1])

		for _, letter := range encrypted_name {
			letters[string(letter)] += 1
		}

		sorted_letters := sort_by_letter_count(letters)

		count := 0
		valid := false
		for _, each := range checksum {
			letter := string(each)
			if sorted_letters[count].Key == string(letter) {
				valid = true
				count += 1
			} else if sorted_letters[count].Value == letters[letter] {
				valid = true
				count += 1
			} else {
				valid = false
				break
			}
		}

		if valid {
			total_sum += sector_id
		}
        fmt.Println(string(room))

	}
	fmt.Println(total_sum)
}
