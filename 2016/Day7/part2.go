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

func find_aba(str string) []string {
    aba_found := []string{}
    for i := 0; i < len(str) - 2; i++ {
        if str[i] != str[i+1] {
            if str[i] == str[i+2] {
                aba_found = append(aba_found, str[i:i+3])
            }
        }
    }
    return aba_found
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
        aba_found, bab_found := []string{}, []string{}

        for i := 0; i < len(str); i++ {
            result := []string{}
            if i % 2 == 0 {
                result = find_aba(str[i])
                for _, each := range result {
                    aba_found = append(aba_found, each)
                }
            } else {
                result = find_aba(str[i])
                for _, each := range result {
                    bab_found = append(bab_found, each)
                }
            }
        }

        for _, aba := range aba_found {
            found := false
            for _, bab := range bab_found {
                a := string(aba)
                b := string(bab)
                if a[0] == b[1] && a[1] == b[0] {
                    valid_ips++
                    found = true
                    break
                }
            }
            if found {
                break
            }
        } 
    }
    fmt.Println(valid_ips)
}
