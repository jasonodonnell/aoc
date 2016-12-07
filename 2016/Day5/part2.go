package main

import (
    "crypto/md5"
    "encoding/hex"
	"flag"
	"fmt"
    "io"
    "strconv"
    "strings"
)

func main() {
	hashPtr := flag.String("md5", "", "MD5 Hash Input")
	flag.Parse()
    
    count, limit := 0, 0
    var password [8]string
    fmt.Println("Decryption In Progress..")

    for limit < 8 {
        h := md5.New()
        io.WriteString(h, *hashPtr + strconv.Itoa(count))
        hash := hex.EncodeToString(h.Sum(nil))

        if strings.Compare(hash[0:5], "00000") == 0 {
            index, err := strconv.Atoi(string(hash[5]))
            if err != nil {
                count++
                continue
            }

            if index < 8 {
                if password[index] == "" { 
                    password[index] = string(hash[6])
                    limit++
                }
            }
        }
        count++
    }

    for _, each := range password { 
        fmt.Print(string(each))
    }

    fmt.Println()
}
