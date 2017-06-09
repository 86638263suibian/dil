package main

import (
    "fmt"
    "bufio"
    "os"
    "flag"
)

func main() {
    flag.Parse()
    args := flag.Args()
    if (len(args) != 2) {
        fmt.Println("wrong usage.")
        return
    }
    current := scan(args[0])
    newS := scan(args[1])
    for _, v := range newS {
        if (! stringInSlice(v, current)) {
            fmt.Println(v)
        }
    }
}

func scan(path string) []string {
    lines := []string{}
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("error")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
