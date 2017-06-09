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
    if (len(args) != 2 || len(args) != 3) {
        fmt.Println("wrong usage.")
        return
    }
    file0 := scan(args[0])
    file1 := scan(args[1])
    for _, v := range file1 {
        if (! stringInSlice(v, file0)) {
            fmt.Println(v)
        }
    }
    // for _, v := range file0 {
    //     if (! stringInSlice(v, file1)) {
    //         fmt.Println(v)
    //     }
    // }
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
