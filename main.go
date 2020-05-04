package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"

    "github.com/joho/godotenv"
)

type Flag struct {
    Name string
    Short string
    Long  string
}

func constructFlag(name, short, long string) Flag {
    var new_flag Flag = Flag{
        Name: name,
        Short: short,
        Long: long,
    }

    return new_flag
}

func getValidFlags() []Flag {
    var verbose Flag = constructFlag("verbose", "-v", "--verbose")

    var flags []Flag = []Flag{verbose}

    return flags 
}

func isFlagInList(flag_name string, flag_list []string) bool {
    for i := 0; i < len(flag_list); i++ {
        if strings.ToLower(flag_name) == strings.ToLower(flag_list[i]) {
            return true
        }
    }

    return false
}

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

//func runList(list_type string, verbose bool) {
func runList(args []string, flags []string) {
    gitPullOrigin(false)
    
    if len(args) != 3 {
        fmt.Println("Please provide a type to list.")
    }
    
    var list_type string = args[2]

    var verbose bool
    if len(flags) > 0 {
        verbose = isFlagInList("verbose", flags)
    }
    
    switch strings.ToLower(list_type) {
        case "books":
            listBooks(verbose)
        case "readinglists":
            listReadingLists(verbose)
        case "users":
            listUsers(verbose)
        default:
            fmt.Printf("Command \"%s\" not found.\n", strings.ToLower(list_type))
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
                printReadingList(&readinglists[reading_list_index], false) 
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

func ValidateFlag(flag string) (is_valid bool, flag_name string) {
    var valid_flags []Flag = getValidFlags()

    for i := 0; i < len(valid_flags); i++ {
        if flag == valid_flags[i].Short || flag == valid_flags[i].Long {
            is_valid = true
            flag_name = valid_flags[i].Name
            return
        }
    }

    return
}

func GetCLIArgs(argslist []string) (args, flags []string) {
    for i := 0; i < len(argslist); i++ {
        is_valid, flag_name := ValidateFlag(argslist[i])
        if is_valid {
            flags = append(flags, flag_name)
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
            if len(os.Args) == 5 {
                value, err := strconv.ParseBool(os.Args[3])
                if err != nil {
                    fmt.Println("Value must be either true or false.")
                }
                
                if os.Args[4] == "--verbose" || os.Args[4] == "-v" {
                    runFilter(os.Args[2], value, true)
                } else {
                    runFilter(os.Args[2], value, false)
                }
            } else if len(os.Args) == 4 {
                value, err := strconv.ParseBool(os.Args[3])
                if err != nil {
                    fmt.Println("Value must be either true or false.")
                }
                
                runFilter(os.Args[2], value, false)
            } else {
                fmt.Println("Please provide a filter and value.")
            }
        case "list":
            runList(args, flags)
            // if len(os.Args) == 3 {
            //     runList(os.Args[2], false)
            // } else if len(os.Args) == 4 {
            //     if os.Args[3] == "--verbose" || os.Args[3] == "-v" {
            //         runList(os.Args[2], true)
            //     }
            // } else {
            //     fmt.Println("Please provide a type to list.")
            // }
        case "readinglist":
            if len(os.Args) == 5 {
                var readinglists []ReadingList = readReadingLists()

                var reading_list_index = getReadingListIndex(os.Args[3])
                if reading_list_index != -1 {
                    if os.Args[4] == "--verbose" || os.Args[4] == "-v" {
                        printReadingList(&readinglists[reading_list_index], true)
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
            if len(os.Args) == 5 {
                if os.Args[4] == "--verbose" || os.Args[4] == "-v" {
                    runSearch(os.Args[2], os.Args[3], true)
                } else {
                    runSearch(os.Args[2], os.Args[3], false)
                }
            } else if len(os.Args) == 4 {
                runSearch(os.Args[2], os.Args[3], false)
            } else {
                fmt.Println("Please provide a search and value.")
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
