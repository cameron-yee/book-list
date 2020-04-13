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
    var books Books
    books = readBooks()
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if my_books[i].Owned == value {
            printBook(my_books[i])
        }
    }
}

func filterRead(value bool) {
    var books Books
    books = readBooks()
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if my_books[i].Read == value {
            printBook(my_books[i])
        }
    }
}

func searchTitle(search_term string) {
    var books Books
    books = readBooks()
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.Contains(strings.ToLower(my_books[i].Title), strings.ToLower(search_term)) {
            printBook(my_books[i])            
        }
    }
}

func searchSeries(search_term string) {
    var books Books
    books = readBooks()
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.Contains(strings.ToLower(my_books[i].Series), strings.ToLower(search_term)) {
            printBook(my_books[i])
        }
    }
}

func searchAuthor(search_term string) {
    var books Books
    books = readBooks()
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.Contains(strings.ToLower(my_books[i].Author), strings.ToLower(search_term)) {
            printBook(my_books[i])            
        }
    }
}

func searchRecommendedBy(search_term string) {
    var books Books
    books = readBooks()
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.Contains(strings.ToLower(my_books[i].RecommendedBy), strings.ToLower(search_term)) {
            printBook(my_books[i])            
        }
    }
}

func searchGenre(search_term string) {
    var books Books
    books = readBooks()
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.Contains(strings.ToLower(my_books[i].Genre), strings.ToLower(search_term)) {
            printBook(my_books[i])            
        }
    }
}
