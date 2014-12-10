package  main

import (
    "fmt"
    "os"
    "bufio"
    "bytes"
)

// Because why Go, whyyyyy?
func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    fmt.Println("Logsplitter v0.2")

    file, err := os.Open("WoWCombatLog.txt")
    check(err)

    defer file.Close()

    var (
        last_date    []byte
        current_date []byte
        outfile      *os.File
    )

    space := []byte(" ")
    slash := []byte("/")
    under := []byte("_")

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        text := scanner.Bytes()
        locn := bytes.Index(text, space)

        current_date = text[:locn]
        if bytes.Compare(current_date, last_date) != 0 {
            fmt.Printf("New date detected: %s\n", current_date)

            last_date = append([]byte{}, current_date...) //force a copy
            if outfile != nil {
                outfile.Close()
            }
            outfile, err = os.Create(
                "wow_log__" +
                string(bytes.Replace(current_date, slash, under, 1)) +
                ".txt")
            check(err)
        }

        outfile.Write(text)
        outfile.WriteString("\n")
    }

    if outfile != nil {
        outfile.Close()
    }

    check(scanner.Err())
}
