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

    fmt.Println("Decryption In Progress..")

    for limit < 8 {
        h := md5.New()
        io.WriteString(h, (*hashPtr + strconv.Itoa(count)))
        hash := hex.EncodeToString(h.Sum(nil))

        if strings.Compare(hash[0:5], "00000") == 0 {
            limit++
            fmt.Print(string(hash[5]))
        }
        count++
    }
    fmt.Println()
}
