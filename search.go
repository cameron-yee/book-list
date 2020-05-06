package main

import (
    "fmt"
    "strconv"
    "strings"
    "os"
)

func runSearch(args []string, flags []Flag) {
    if len(args) != 4 {
        fmt.Println("Please provide a field and value.")
    }
    
    var verbose bool
    var vverbose bool
    var limit int
    
    if len(flags) > 0 {
        var verbose_flag *Flag = GetFlag("verbose", flags)
        var limit_flag *Flag = GetFlag("limit", flags)

        if verbose_flag != nil {
            verbose = true
        }
        
        if len(flags) > 0 {
            var vverbose_flag *Flag = GetFlag("vverbose", flags)

            if vverbose_flag != nil {
                vverbose = true
            }
        }

        if limit_flag != nil {
            i_value, err := strconv.ParseInt(GetFlagValue(limit_flag), 10, 0)
            if err != nil {
                fmt.Println("Limit value must be an integer.")
                fmt.Printf("%v\n", err)
                os.Exit(1)
            }

            limit = int(i_value)
        }
    }
    
    var label string = args[2]
    var search_term string = args[3]
    switch strings.ToLower(label) {
        case "title":
            searchTitle(search_term, verbose, vverbose, limit)
        case "series":
            searchSeries(search_term, verbose, vverbose, limit)
        case "author":
            searchAuthor(search_term, verbose, vverbose, limit)
        case "recommendedby":
            searchRecommendedBy(search_term, verbose, vverbose, limit)
        case "genre":
            searchGenre(search_term, verbose, vverbose, limit)
        case "readby":
            searchReadBy(search_term, verbose, vverbose, limit)
        default:
            fmt.Println("Please provide a search and value.")
    }
}

func runFilter(args []string, flags []Flag) {
    if len(args) != 4 {
        fmt.Println("Please provide a filter and value.")
    }

    var verbose bool
    var vverbose bool
    var limit int
    
    if len(flags) > 0 {
        var verbose_flag *Flag = GetFlag("verbose", flags)
        var limit_flag *Flag = GetFlag("limit", flags)

        if verbose_flag != nil {
            verbose = true
        }
        
        if len(flags) > 0 {
            var vverbose_flag *Flag = GetFlag("vverbose", flags)

            if vverbose_flag != nil {
                vverbose = true
            }
        }

        if limit_flag != nil {
            i_value, err := strconv.ParseInt(GetFlagValue(limit_flag), 10, 0)
            if err != nil {
                fmt.Println("Limit value must be an integer.")
                fmt.Printf("%v\n", err)
                os.Exit(1)
            }

            limit = int(i_value)
        }
    }
    
    var label string = args[2]
    var value string = args[3]
    value_as_bool, err := strconv.ParseBool(value)
    if err != nil {
        fmt.Println("Value must be either true or false.")
        return
    }
    
    switch strings.ToLower(label) {
        case "owned":
            filterOwned(value_as_bool, verbose, vverbose, limit)
        default:
            fmt.Println("The only filter allowed is \"owned\"")
    }
}

func filterOwned(value, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var until int = len(books)
    if limit != 0 {
        until = limit
    }

    for i := 0; i < until; i++ {
        if books[i].Owned == value {
            printBook(&books[i], false, verbose, vverbose)
        }
    }
}

func searchReadBy(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var until int = len(books)
    if limit != 0 {
        until = limit
    }

    for i := 0; i < until; i++ {
        for j := 0; j < len(books[i].ReadBy); j++ {
            if containsCaseInsensitive(books[i].ReadBy[j], search_term) {
                printBook(&books[i], false, verbose, vverbose)
            }
        }
    }
}

func searchTitle(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var until int = len(books)
    if limit != 0 {
        until = limit
    }

    for i := 0; i < until; i++ {
        if containsCaseInsensitive(books[i].Title, search_term) {
            printBook(&books[i], false, verbose, vverbose)
        }
    }
}

func searchSeries(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var until int = len(books)
    if limit != 0 {
        until = limit
    }

    for i := 0; i < until; i++ {
        if containsCaseInsensitive(books[i].Series, search_term) {
            printBook(&books[i], false, verbose, vverbose)
        }
    }
}

func searchAuthor(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var until int = len(books)
    if limit != 0 {
        until = limit
    }

    for i := 0; i < until; i++ {
        if containsCaseInsensitive(books[i].Author, search_term) {
            printBook(&books[i], false, verbose, vverbose)            
        }
    }
}

func searchRecommendedBy(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var until int = len(books)
    if limit != 0 {
        until = limit
    }

    for i := 0; i < until; i++ {
        if containsCaseInsensitive(books[i].RecommendedBy, search_term) {
            printBook(&books[i], false, verbose, vverbose)
        }
    }
}

func searchGenre(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var until int = len(books)
    if limit != 0 {
        until = limit
    }

    for i := 0; i < until; i++ {
        if containsCaseInsensitive(books[i].Genre, search_term) {
            printBook(&books[i], false, verbose, vverbose)
        }
    }
}
