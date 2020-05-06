package main

import (
    "fmt"
    "strconv"
    "strings"
)

func runSearch(args []string, flags []Flag) {
    if len(args) != 4 {
        fmt.Println("Please provide a field and value.")
    }
    
    var label string = args[2]
    var search_term string = args[3]
    
    var verbose bool = getExistsFlagValue("verbose", &flags)
    var vverbose bool = getExistsFlagValue("vverbose", &flags)
    var limit int = getStoreIntFlagValue("limit", &flags)
    
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
    
    var label string = args[2]
    var value string = args[3]
    
    var verbose bool = getExistsFlagValue("verbose", &flags)
    var vverbose bool = getExistsFlagValue("vverbose", &flags)
    var limit int = getStoreIntFlagValue("limit", &flags)

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
    
    var count int = 0
    for i := 0; i < len(books); i++ {
        if books[i].Owned == value {
            printBook(&books[i], false, verbose, vverbose)
            count++
        }

        if limit != 0 && count == limit {
            break
        }
    }
}

func searchReadBy(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var count int = 0

Loops:
    for i := 0; i < len(books); i++ {
        for j := 0; j < len(books[i].ReadBy); j++ {
            if containsCaseInsensitive(books[i].ReadBy[j], search_term) {
                printBook(&books[i], false, verbose, vverbose)
                count++

                if limit != 0 && count == limit {
                    break Loops
                }
            }
        }
    }
}

func searchTitle(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var count int = 0
    for i := 0; i < len(books); i++ {
        if containsCaseInsensitive(books[i].Title, search_term) {
            printBook(&books[i], false, verbose, vverbose)
            count++
        }

        if limit != 0 && count == limit {
            break
        }
    }
}

func searchSeries(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var count int = 0
    for i := 0; i < len(books); i++ {
        if containsCaseInsensitive(books[i].Series, search_term) {
            printBook(&books[i], false, verbose, vverbose)
            count++
        }

        if limit != 0 && count == limit {
            break
        }
    }
}

func searchAuthor(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var count int = 0
    for i := 0; i < len(books); i++ {
        if containsCaseInsensitive(books[i].Author, search_term) {
            printBook(&books[i], false, verbose, vverbose)
            count++
        }

        if limit != 0 && count == limit {
            break
        }
    }
}

func searchRecommendedBy(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var count int = 0
    for i := 0; i < len(books); i++ {
        if containsCaseInsensitive(books[i].RecommendedBy, search_term) {
            printBook(&books[i], false, verbose, vverbose)
            count++
        }

        if limit != 0 && count == limit {
            break
        }
    }
}

func searchGenre(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var count int = 0
    for i := 0; i < len(books); i++ {
        if containsCaseInsensitive(books[i].Genre, search_term) {
            printBook(&books[i], false, verbose, vverbose)
            count++
        }

        if limit != 0 && count == limit {
            break
        }
    }
}
