package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func help() {
    fmt.Println("Commands:")
    fmt.Println("\tadd: Add a book, readinglist, or user. Additionally, add a member to a reading list.")
    fmt.Println("\tdelete: Delete a book, readinglist, or user. Additionally, delete a member from a reading list.")
    fmt.Println("\tfilter: Filter based on bool property. Ex. ./main.go filter owned true")
    fmt.Println("\thelp: Print list of commands.")
    fmt.Println("\tlist: List books, readinglists, or users.")
    fmt.Println("\tsearch: Search based on string property. Ex. ./main.go search title Narnia")
    fmt.Println("\tupdate: Update a type by providing a type, primary key, field, and the new value for the field.")
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

func runList(list_type string) {
    switch strings.ToLower(list_type) {
        case "readinglists":
            listReadingLists()
        case "users":
            listUsers()
        default:
            listBooks()
    }
}

func runUpdate(update_type, update_primary_key, update_field, update_field_value string) {
    switch strings.ToLower(update_type) {
        case "readinglist":
            runUpdateReadingList(update_primary_key, update_field, update_field_value)
        case "user":
            runUpdateUser(update_primary_key, update_field, update_field_value)
        default:
            runUpdateBook(update_primary_key, update_field, update_field_value)
    }
}

func main() {
    if len(os.Args) < 2 {
        help()
        return
    }

    switch command := os.Args[1]; command {
        case "add":
            if len(os.Args) == 3 {
                runAdd(os.Args[2])
            } else if len(os.Args) == 5 {
                updateReadingListAddMember(os.Args[3], os.Args[4])
            } else {
                panic("Please provide a type to add")
            }
        case "delete":
            if len(os.Args) == 4 {
                runDelete(os.Args[2], os.Args[3])
            } else if (len(os.Args) == 5) {
                updateReadingListDeleteMember(os.Args[3], os.Args[4])
            } else {
                panic("Please provide a type to delete and a value")
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
            if len(os.Args) != 3 {
                panic("Please provide a type to list")
            } else {
                runList(os.Args[2])
            }
        case "search":
            if len(os.Args) != 4 {
                panic("Please provide a search and value.")
            } else {
                runSearch(os.Args[2], os.Args[3])
            }
        case "update":
            if len(os.Args) != 6 {
                panic("Please provide a type, primary_key, field to update, and the new value for the field.") // readinglist, "main", title, "main2"
            } else {
                runUpdate(os.Args[2], os.Args[3], os.Args[4], os.Args[5]) // title, field, field_value
            }
        default:
            help()
    }
}
