package main

import (
    "strings"
)


func runSearch(label, search_term string) {
    switch strings.ToLower(label) {
        case "series":
            searchSeries(search_term)
        case "author":
            searchAuthor(search_term)
        case "recommendedby":
            searchRecommendedBy(search_term)
        case "genre":
            searchGenre(search_term)
        default:
            searchTitle(search_term)
    }
}

func runFilter(label string, value bool) {
    switch strings.ToLower(label) {
        case "owned":
            filterOwned(value)
        default:
            filterRead(value)
    }
}

func filterOwned(value bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if books[i].Owned == value {
            printBook(books[i], false, false)
        }
    }
}

func filterRead(value bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if books[i].Read == value {
            printBook(books[i], false, false)
        }
    }
}

func searchTitle(search_term string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].Title), strings.ToLower(search_term)) {
            printBook(books[i], false, false)            
        }
    }
}

func searchSeries(search_term string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].Series), strings.ToLower(search_term)) {
            printBook(books[i], false, false)
        }
    }
}

func searchAuthor(search_term string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].Author), strings.ToLower(search_term)) {
            printBook(books[i], false, false)            
        }
    }
}

func searchRecommendedBy(search_term string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].RecommendedBy), strings.ToLower(search_term)) {
            printBook(books[i], false, false)            
        }
    }
}

func searchGenre(search_term string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].Genre), strings.ToLower(search_term)) {
            printBook(books[i], false, false)            
        }
    }
}
