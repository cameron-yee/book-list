package main

import (
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/joho/godotenv"
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
        case "book":
            addBook()
        case "readinglist":
            addReadingList()
        case "user":
            addUser()
        default:
            fmt.Printf("Command \"%s\" not found.\n", strings.ToLower(add_type))
    }
}

func runDelete(delete_type, value string) {
    switch strings.ToLower(delete_type) {
        case "book":
            deleteBook(value)
        case "readinglist":
            deleteReadingList(value)
        case "user":
            deleteUser(value)
        default:
            fmt.Printf("Command \"%s\" not found.\n", strings.ToLower(delete_type))
    }
}

func runList(args []string, flags []Flag) {
    gitPullOrigin(false)
    
    if len(args) != 3 {
        fmt.Println("Please provide a type to list.")
    }
    
    var list_type string = args[2]
    
    var verbose bool = getExistsFlagValue("verbose", &flags)
    var vverbose bool = getExistsFlagValue("vverbose", &flags)
    var limit int = getStoreIntFlagValue("limit", &flags)

    switch strings.ToLower(list_type) {
        case "books":
            listBooks(verbose, vverbose, limit)
        case "readinglists":
            listReadingLists(verbose, vverbose, limit)
        case "users":
            listUsers(verbose, limit)
        default:
            fmt.Printf("Command \"%s\" not found.\n", strings.ToLower(list_type))
    }
}

func runReadingList(args []string, flags []Flag) {
    if len(args) != 4 {
        fmt.Println("Please provide a type to list.")
    }
    
    var command string = args[2]
    var value string = args[3]

    var verbose bool = getExistsFlagValue("verbose", &flags)
    var vverbose bool = getExistsFlagValue("vverbose", &flags)
    
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
                printReadingList(&readinglists[reading_list_index], verbose, vverbose) 
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

func GetCLIArgs(argslist []string) (args []string, flags []Flag) {
    for i := 0; i < len(argslist); i++ {
        is_valid, flag := ValidateFlag(argslist[i])
        
        if is_valid {
            if (flag.Action == "store") {
                if i == len(argslist) - 1 {
                    fmt.Printf("Flag \"%v\" requires a value.\n", flag.Name)
                }
                
                flag.Value = argslist[i+1]
                i++
            }
            
            flags = append(flags, flag)
        } else {
            args = append(args, argslist[i])
        } 
    }

    return
}

func init() {
    var pathname string = getCallDirectory() 
    
    if err := godotenv.Load(fmt.Sprintf("%s/.env", pathname)); err != nil {
        log.Fatal("no .env file found")
    }
}

func main() {
    if len(os.Args) < 3 {
        help()
        return
    }

    args, flags := GetCLIArgs(os.Args)

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
            runFilter(args, flags)
        case "list":
            runList(args, flags)
        case "readinglist":
            runReadingList(args, flags)
        case "search":
            runSearch(args, flags)
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
