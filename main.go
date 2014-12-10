package  main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "bytes"
    // "log"
)

// Because why Go, whyyyyy?
func check(err error) {
    if err != nil {
        // log.Fatal(err)
        panic(err)
    }
}

func main() {
    fmt.Println("Logsplitter v0.2")

    file, err := os.Open("WoWCombatLog.txt")
    check(err)

    defer file.Close()

    var last_date string
    var current_date string
    var outfile *os.File

    space := []byte(" ")

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        text := scanner.Bytes()
        locn := bytes.Index(text, space)

        current_date = string(text[:locn])
        // if bytes.Compare(current_date, last_date) != 0 {
        if current_date != last_date {
            fmt.Println("New date detected:", current_date)

            last_date = current_date
            if outfile != nil {
                outfile.Close()
            }
            outfile, err = os.Create("wow_log__" + strings.Replace(current_date, "/", "_", 1) + ".txt")
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
