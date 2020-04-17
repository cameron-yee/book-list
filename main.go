package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func help() {
    fmt.Println("Commands:")
    fmt.Println("\tadd: Add a book.")
    fmt.Println("\tdelete: Delete a book by supplying the book title.")
    fmt.Println("\tfilter: Filter based on bool property. Ex. ./main.go filter owned true")
    fmt.Println("\thelp: Print list of commands.")
    fmt.Println("\tlist: List books.")
    fmt.Println("\tsearch: Search based on string property. Ex. ./main.go search title Narnia")
    fmt.Println("\tupdate: Update a book by providing a book title, a field, and the new value for the field.")
}

func runAdd(add_type string) {
    switch strings.ToLower(add_type) {
        case "readinglist":
            addReadingList()
        case "user":
            addUser()
        default:
            addBook()
    }
}

func runDelete(delete_type, value string) {
    switch strings.ToLower(delete_type) {
        case "readinglist":
            deleteReadingList(value)
        case "user":
            deleteUser(value)
        default:
            deleteBook(value)
    }
}

func main() {
    if len(os.Args) < 2 {
        help()
        return
    }

    switch command := os.Args[1]; command {
        case "add":
            if len(os.Args) != 3 {
                panic("Please provide a type to add")
            } else {
                runAdd(os.Args[2])
            }
        case "delete":
            if len(os.Args) != 4 {
                panic("Please provide a type to delete and a value")
            } else {
                runDelete(os.Args[2], os.Args[3])
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
        case "list":
            listBooks()
        case "search":
            if len(os.Args) != 4 {
                panic("Please provide a search and value.")
            } else {
                runSearch(os.Args[2], os.Args[3])
            }
        case "update":
            if len(os.Args) != 5 {
                panic("Please provide a title, a field to update, and the new value for the field.")
            } else {
                runUpdate(os.Args[2], os.Args[3], os.Args[4]) // title, field, field_value
            }
        default:
            help()
    }
}
