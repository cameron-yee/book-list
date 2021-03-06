package main

import (
    "fmt"
    "strconv"
    "strings"
)

func runSearch(args []string, flags FlagList) {
    var help bool = GetBoolFlagValue("help", flags)
    
    if help {
        fmt.Println("Description:")
        fmt.Println("\tSearch for books based on field value.")
        fmt.Println("Usage:")
        fmt.Println("\tbooklist search <-h, -v, -vv, -l> <title, series, author, recommendedby, genre, readby> [search_value]")
        return
    }
    
    if len(args) != 4 {
        fmt.Println("Please provide a field and value.")
        return
    }
    
    var label string = args[2]
    var search_term string = args[3]
    
    var falsee bool = GetBoolFlagValue("false", flags)
    var verbose bool = GetBoolFlagValue("verbose", flags)
    var vverbose bool = GetBoolFlagValue("vverbose", flags)
    var limit int = GetIntFlagValue("limit", flags)
    
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
            searchReadBy(search_term, verbose, vverbose, limit, falsee)
        default:
            fmt.Println("Please provide a search and value.")
    }
}

func runFilter(args []string, flags FlagList) {
    var help bool = GetBoolFlagValue("help", flags)
    
    if help {
        fmt.Println("Description:")
        fmt.Println("\tFilter books based on field value.")
        fmt.Println("Usage:")
        fmt.Println("\tbooklist filter <-h, -v, -vv, -l> <owned> [true, false]")
        return
    }
    
    if len(args) != 4 {
        fmt.Println("Please provide a filter and value.")
        return
    }
    
    var label string = args[2]
    var value string = args[3]
    
    var verbose bool = GetBoolFlagValue("verbose", flags)
    var vverbose bool = GetBoolFlagValue("vverbose", flags)
    var limit int = GetIntFlagValue("limit", flags)

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

func searchReadBy(search_term string, verbose bool, vverbose bool, limit int, falsee bool) {
    var books []Book = readBooks()
    
    var count int = 0
    
    var booksDontContainSearchTerm []int

Loops:
    for i := 0; i < len(books); i++ {
        var bookDoesntContainSearchTerm bool = true
        
        for j := 0; j < len(books[i].ReadBy); j++ {
            if falsee && containsCaseInsensitive(books[i].ReadBy[j], search_term) {
                bookDoesntContainSearchTerm = false
            }
            
            if !falsee && containsCaseInsensitive(books[i].ReadBy[j], search_term) {
                printBook(&books[i], false, verbose, vverbose)
                count++
            }
            
            if limit != 0 && count == limit {
                break Loops
            }
        }

        if bookDoesntContainSearchTerm {
            booksDontContainSearchTerm = append(booksDontContainSearchTerm, i)
        }
    }

    if falsee {
        for i := 0; i < len(booksDontContainSearchTerm); i++ { 
            printBook(&books[booksDontContainSearchTerm[i]], false, verbose, vverbose)
            count ++
            
            if limit != 0 && count == limit {
                break
            }
        }
    }
}

func searchTitle(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var searchResultsFound bool
    var count int = 0
    
    for i := 0; i < len(books); i++ {
        if containsCaseInsensitive(books[i].Title, search_term) {
            printBook(&books[i], false, verbose, vverbose)
            searchResultsFound = true
            count++
        }

        if limit != 0 && count == limit {
            break
        }
    }

    if !searchResultsFound {
        fmt.Printf("No book found with title: %s\n", search_term)
    }
}

func searchSeries(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var searchResultsFound bool
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
    
    if !searchResultsFound {
        fmt.Printf("No book found with title: %s\n", search_term)
    }
}

func searchAuthor(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var searchResultsFound bool
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
    
    if !searchResultsFound {
        fmt.Printf("No book found with author: %s\n", search_term)
    }
}

func searchRecommendedBy(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()
    
    var searchResultsFound bool
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
    
    if !searchResultsFound {
        fmt.Printf("No book recommended by: %s\n", search_term)
    }
}

func searchGenre(search_term string, verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()

    var searchResultsFound bool
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
    
    if !searchResultsFound {
        fmt.Printf("No book found with genre: %s\n", search_term)
    }
}
