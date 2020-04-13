package main

import (
    "fmt"
    "os"
    "strconv"
)

func help() {
    fmt.Println("Commands:")
    fmt.Println("\thelp: Print list of commands.")
    fmt.Println("\tlist: List books.")
    fmt.Println("\tadd: Add a book.")
    fmt.Println("\tfilter: Filter based on bool property. Ex. ./main.go filter owned true")
    fmt.Println("\tsearch: Search based on string property. Ex. ./main.go search title Narnia")
}

func main() {
    if len(os.Args) < 2 {
        help()
        return
    }

    switch command := os.Args[1]; command {
        case "list":
            listBooks()
        case "add":
            addBook()
        case "search":
            if len(os.Args) != 4 {
                panic("Please provide a search and value.")
            } else {
                runSearch(os.Args[2], os.Args[3])
            }
        case "filter":
            if len(os.Args) != 4 {
                panic("Please provide a filter and value.")
            } else {
                value, err := strconv.ParseBool(os.Args[3])
                if err != nil {
                    panic(err)
                }
                
                runFilter(os.Args[2], value)
            } 
        default:
            help()
    }
}
