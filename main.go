package  main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
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
    fmt.Println("Logsplitter, starting up.")

    file, err := os.Open("infile.txt")
    check(err)

    defer file.Close()

    var last_date string
    var current_date string
    var outfile *os.File

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        s := strings.Split(scanner.Text(), " ")

        current_date = s[0]
        if current_date != last_date {
            fmt.Println("New date detected:", current_date)
            last_date = current_date
            if outfile != nil {
                outfile.Close()
            }
            outfile, err = os.Create("wow_log__" + strings.Replace(current_date, "/", "_", 1) + ".txt")
            check(err)
        }

        outfile.Write(scanner.Bytes())
    }

    if outfile != nil {
        outfile.Close()
    }

    check(scanner.Err())

}
