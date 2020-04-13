package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
)

type Book struct {
    Title         string  `json:"title"`
    Series        string  `json:"series"`
    Author        string  `json:"author"`
    RecommendedBy string  `json:"recommendedBy"`
    Read          bool    `json:"read"`
    Owned         bool    `json:"owned"`
    Genre         string  `json:"genre"`
}

type Books struct {
    Books []Book `json:"books"`
}

func readBooks() Books {
    data, err := ioutil.ReadFile("./books.json")
    if err != nil {
        panic(err)
    }


    var books Books

    err = json.Unmarshal(data, &books)
    if err != nil {
        panic(err)
    }

    return books
}

func printBook(book Book) {
    colorPrintString("Title", book.Title)
    colorPrintString("Series", book.Series)
    colorPrintString("Author", book.Author)
    colorPrintString("Recommended By", book.RecommendedBy)
    colorPrintBool("Read", book.Read)
    colorPrintBool("Owned", book.Owned)
    colorPrintString("Genre", book.Genre)
    fmt.Println("-------------------------------------------------------------")
}

func listBooks() {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        printBook(my_books[i])
    }
}

func writeBook(book *Book) {
    var books Books
    books = readBooks()
    
    books.Books = append(books.Books, *book)

    dataBytes, err := json.Marshal(books)
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile("./books.json", dataBytes, 0644)
    if err != nil {
        panic(err)
    }
}

func addBook() {
    title := getInput("Title")
    series := getInput("Series")
    author := getInput("Author")
    recommended_by := getInput("Recommended By")
    
    read := getInput("Read")
    read_bool := false
    if strings.ToLower(read) == "true" {
        read_bool = true    
    }
    
    owned := getInput("Owned")
    owned_bool := false
    if strings.ToLower(owned) == "true" {
        owned_bool = true    
    }
    
    genre := getInput("Genre")

    var new_book *Book = &Book{
        Title: title,
        Series: series,
        Author: author,
        RecommendedBy: recommended_by,
        Read: read_bool,
        Owned: owned_bool,
        Genre: genre,
    }

    writeBook(new_book)
}


