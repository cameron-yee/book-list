package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"

    "github.com/joho/godotenv"
)

func help() {
    fmt.Println("Usage:")
    fmt.Println("\tbooklist [command] <flags> <options>")
    fmt.Println("Commands:")
    fmt.Println("\tadd [book, readinglist, user]")
    fmt.Println("\tdelete [book, readinglist, user]")
    fmt.Println("\tfilter [owned]")
    fmt.Println("\thelp")
    fmt.Println("\tlist [books, readinglists, users]")
    fmt.Println("\tsearch [title, series, author, recommendedby, genre, readby] <search-value>")
    fmt.Println("\tupdate [book, readinglist, user]")
    fmt.Println("\treadinglist [add, delete, print] <-v, -vv, -l> <book, member>")
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
            fmt.Printf("Command \"%s\" not found. Options are book, readinglist, or user.\n", strings.ToLower(add_type))
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
            fmt.Printf("Command \"%s\" not found. Options are book, readinglist, or user.\n", strings.ToLower(delete_type))
    }
}

func runList(args []string, flags FlagList) {
    gitPullOrigin(false)

    var help bool = GetBoolFlagValue("help", flags)
    
    if help {
        fmt.Println("Description:")
        fmt.Println("\tList values for saved books, readinglists, or users.")
        fmt.Println("Usage:")
        fmt.Println("\tbooklist list <-h> [books, readinglists, users]")
        return
    }

    if len(args) != 3 {
        fmt.Println("Please provide a type to list.")
    }
    
    var list_type string = args[2]

    var verbose bool = GetBoolFlagValue("verbose", flags)
    var vverbose bool = GetBoolFlagValue("vverbose", flags)
    var limit int = GetIntFlagValue("limit", flags)

    switch strings.ToLower(list_type) {
        case "books":
            listBooks(verbose, vverbose, limit)
        case "readinglists":
            listReadingLists(verbose, vverbose, limit)
        case "users":
            listUsers(verbose, limit)
        default:
            fmt.Printf("Command \"%s\" not found. Options are books, readinglists, or users.\n", strings.ToLower(list_type))
    }
}

func runReadingList(args []string, flags FlagList) {
    if len(args) != 4 {
        fmt.Println("Please provide a type to list.")
    }
    
    var command string = args[2]
    var value string = args[3]

    var verbose bool = GetBoolFlagValue("verbose", flags)
    var vverbose bool = GetBoolFlagValue("vverbose", flags)
    
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
                return
            }
            
            fmt.Printf("No readinglist with title: \"%s\".\n", value)
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
            fmt.Printf("Command \"%s\" not found. Options are book, readinglist, or user.", strings.ToLower(update_type))
    }
}

func checkForFlagValue(argslist *[]string, i int) {
    if i == len(*argslist) - 1 {
        fmt.Printf("Flag \"%s\" requires a value.\n", (*argslist)[i])
        os.Exit(1)
    }
}

func GetCLIArgs(argslist []string) ([]string, FlagList) {
    var args []string
    var validFlags FlagList = getValidFlags()
        
    for i := 0; i < len(argslist); i++ {
        var flagType string = GetFlagType(argslist[i], validFlags)

        switch (flagType) {
            case "":
                args = append(args, argslist[i])
            case "bool":
                for j := 0; j < len(validFlags.BoolFlags); j++ {
                    if validFlags.BoolFlags[j].Flag.Short == argslist[i] ||
                       validFlags.BoolFlags[j].Flag.Long == argslist[i] {
                        validFlags.BoolFlags[j].Value = true
                    }
                }
            case "int":
                checkForFlagValue(&argslist, i)
                
                for j := 0; j < len(validFlags.IntFlags); j++ {
                    if validFlags.IntFlags[j].Flag.Short == argslist[i] ||
                       validFlags.IntFlags[j].Flag.Long == argslist[i] {
                        value_to_int, err := strconv.Atoi(argslist[i+1])
                        if err != nil {
                            fmt.Printf("Value for %s must be an integer.", validFlags.IntFlags[j].Flag.Name)
                        }
                        
                        validFlags.IntFlags[j].Value = value_to_int 
                    }
                }

                i++
            case "string":
                checkForFlagValue(&argslist, i)
                
                for j := 0; j < len(validFlags.StringFlags); j++ {
                    if validFlags.StringFlags[j].Flag.Short == argslist[i] ||
                       validFlags.StringFlags[j].Flag.Long == argslist[i] {
                        validFlags.StringFlags[j].Value = argslist[i+1]
                    }
                }

                i++
        }
    }

    return args, validFlags
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
                return
            }
            
            fmt.Println("Please provide a type to add.")
        case "delete":
            if len(os.Args) == 4 {
                runDelete(os.Args[2], os.Args[3])
                return
            }
            
            fmt.Println("Please provide a type to delete and a value.")
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
                return
            }
            
            runUpdate(os.Args[2]) // title, field, field_value
        default:
            help()
    }
}
