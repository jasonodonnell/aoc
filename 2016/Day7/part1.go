package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

func find_abba(str string) bool {
    for i := 0; i < len(str) - 3; i++ {
        if str[i] != str[i+1] {
            if str[i] == str[i+3] && str[i+1] == str[i+2] {
                return true
            }
        }
    }
    return false
}

func main() {
	filenamePtr := flag.String("filename", "", "Input data filename")
	flag.Parse()

	ips := read_file(*filenamePtr)
    valid_ips := 0

    for _, ip := range ips {
        ip = strings.Replace(ip, "[", "-", -1)
        ip = strings.Replace(ip, "]", "-", -1)
        str := strings.Split(ip, "-")
        hyper_valid := false
        ip_valid := false

        for i := 0; i < len(str); i++ {
            if i % 2 == 0 {
                if find_abba(str[i]) {
                    ip_valid = true
                }
            } else {
                if find_abba(str[i]) {
                    hyper_valid = false
                    break
                } else {
                    hyper_valid = true
                }
            }
        }

        if ip_valid && hyper_valid {
            valid_ips++
        }
    }
    fmt.Println(valid_ips)
}
