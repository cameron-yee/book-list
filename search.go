package main

import (
    "strings"
)


func runSearch(label, search_term string, verbose bool) {
    switch strings.ToLower(label) {
        case "series":
            searchSeries(search_term, verbose)
        case "author":
            searchAuthor(search_term, verbose)
        case "recommendedby":
            searchRecommendedBy(search_term, verbose)
        case "genre":
            searchGenre(search_term, verbose)
        default:
            searchTitle(search_term, verbose)
    }
}

func runFilter(label string, value bool, verbose bool) {
    switch strings.ToLower(label) {
        case "owned":
            filterOwned(value, verbose)
        default:
            filterRead(value, verbose)
    }
}

func filterOwned(value, verbose bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if books[i].Owned == value {
            printBook(books[i], false, verbose)
        }
    }
}

func filterRead(value, verbose bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if books[i].Read == value {
            printBook(books[i], false, verbose)
        }
    }
}

func searchTitle(search_term string, verbose bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].Title), strings.ToLower(search_term)) {
            printBook(books[i], false, verbose)
        }
    }
}

func searchSeries(search_term string, verbose bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].Series), strings.ToLower(search_term)) {
            printBook(books[i], false, verbose)
        }
    }
}

func searchAuthor(search_term string, verbose bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].Author), strings.ToLower(search_term)) {
            printBook(books[i], false, verbose)            
        }
    }
}

func searchRecommendedBy(search_term string, verbose bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].RecommendedBy), strings.ToLower(search_term)) {
            printBook(books[i], false, verbose)
        }
    }
}

func searchGenre(search_term string, verbose bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.Contains(strings.ToLower(books[i].Genre), strings.ToLower(search_term)) {
            printBook(books[i], false, verbose)
        }
    }
}
