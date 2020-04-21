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

func runList(list_type string, verbose bool) {
    switch strings.ToLower(list_type) {
        case "readinglists":
            listReadingLists(verbose)
        case "users":
            listUsers()
        default:
            listBooks(verbose)
    }
}

func runReadingList(command, value string) {
    switch strings.ToLower(command) {
        case "add":
            if strings.ToLower(value) == "book" {
                addBookToReadingList()
            } else if strings.ToLower(value) == "member" {
                addMemberToReadingList()
            }
        case "delete":
            if strings.ToLower(value) == "book" {
                deleteBookFromReadingList()
            } else if strings.ToLower(value) == "member" {
                deleteMemberFromReadingList()
            }
        case "print":
            var reading_list_index int = getReadingListIndex(value)

            if reading_list_index != -1 {
                var readinglists []ReadingList = readReadingLists()
                printReadingList(readinglists[reading_list_index], false) 
            } else {
                fmt.Printf("No readinglist with title: \"%s\".\n", value)
            }
        default:
            fmt.Println("Options are add, delete, or print.")
    }
}

func runUpdate(update_type string) {
    switch strings.ToLower(update_type) {
        case "book":
            runUpdateBook()
        case "readinglist":
            runUpdateReadingList()
        case "user":
            runUpdateUser()
        default:
            fmt.Println("Options are book, readinglist, or user.")
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
            } else {
                fmt.Println("Please provide a type to add.")
            }
        case "delete":
            if len(os.Args) == 4 {
                runDelete(os.Args[2], os.Args[3])
            } else {
                fmt.Println("Please provide a type to delete and a value.")
            }
        case "filter":
            if len(os.Args) != 4 {
                fmt.Println("Please provide a filter and value.")
            } else {
                value, err := strconv.ParseBool(os.Args[3])
                if err != nil {
                    panic(err)
                }
                
                runFilter(os.Args[2], value)
            }
        case "list":
            if len(os.Args) == 3 {
                runList(os.Args[2], false)
            } else if len(os.Args) == 4 {
                if os.Args[3] == "--verbose" || os.Args[3] == "-v" {
                    runList(os.Args[2], true)
                }
            } else {
                fmt.Println("Please provide a type to list.")
            }
        case "readinglist":
            if len(os.Args) == 5 {
                var readinglists []ReadingList = readReadingLists()

                var reading_list_index = getReadingListIndex(os.Args[3])
                if reading_list_index != -1 {
                    if os.Args[4] == "--verbose" || os.Args[4] == "-v" {
                        printReadingList(readinglists[reading_list_index], true)
                    }
                } else {
                    fmt.Printf("No readinglist with title: \"%s\".\n", os.Args[3])
                }
            } else if len(os.Args) == 4 {
                runReadingList(os.Args[2], os.Args[3])
            } else {
                fmt.Println("add/delete book or member. print readinglist.")
            }
        case "search":
            if len(os.Args) != 4 {
                fmt.Println("Please provide a search and value.")
            } else {
                runSearch(os.Args[2], os.Args[3])
            }
        case "update":
            if len(os.Args) != 3 {
                fmt.Println("Please provide a type to update.") // readinglist, "main", title, "main2"
            } else {
                runUpdate(os.Args[2]) // title, field, field_value
            }
        default:
            help()
    }
}
