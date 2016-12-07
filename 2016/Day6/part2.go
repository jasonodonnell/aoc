package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
    "sort"
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
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	signals := read_file(*filenamePtr)

    const limit = 8 
    count := 0

    for count < limit {
        letters := make(map[string]int)

        for _, letter := range signals {
            letters[string(letter[count])] += 1
        }

        sorted_letters := sort_by_letter_count(letters)
        fmt.Printf("%s", sorted_letters[0].Key)
        count++
    }
    fmt.Println("")
}
