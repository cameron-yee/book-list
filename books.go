package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

type Book struct {
    Title         string    `json:"title"`
    Series        string    `json:"series"`
    Author        string    `json:"author"`
    ReadingList  string    `json:"readingList"`
    RecommendedBy string    `json:"recommendedBy"`
    Read          bool      `json:"read"`
    Owned         bool      `json:"owned"`
    Genre         string    `json:"genre"`
    EntryOwner    string    `json:"entryOwner"`
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
    colorPrintString("ReadingList", book.ReadingList)
    colorPrintString("Recommended By", book.RecommendedBy)
    colorPrintBool("Read", book.Read)
    colorPrintBool("Owned", book.Owned)
    colorPrintString("Genre", book.Genre)
    colorPrintString("EntryOwner", book.EntryOwner)
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

func writeBooks(books Books) {
    dataBytes, err := json.Marshal(books)
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile("./books.json", dataBytes, 0644)
    if err != nil {
        panic(err)
    }
}

func appendBook(book *Book) {
    var books Books
    books = readBooks()
    
    books.Books = append(books.Books, *book)

    writeBooks(books)
}

func addBook() {
    title := getInput("Title")
    series := getInput("Series")
    author := getInput("Author")
    reading_list := getInput("Reading List")
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
    entry_owner := getInput("Entry Owner")

    var new_book *Book = &Book{
        Title: title,
        Series: series,
        Author: author,
        ReadingList: reading_list,
        RecommendedBy: recommended_by,
        Read: read_bool,
        Owned: owned_bool,
        Genre: genre,
        EntryOwner: entry_owner,
    }

    appendBook(new_book)
}

func runUpdate(book_title, field, value string) {
    switch strings.ToLower(field) {
        case "title":
            updateBookTitle(book_title, value)
        case "series":
            updateBookSeries(book_title, value)
        case "author":
            updateBookAuthor(book_title, value)
        case "readinglist":
            updateBookReadingList(book_title, value)
        case "recommendedby":
            updateBookRecommendedBy(book_title, value)
        case "read":
            value_as_bool, _ := strconv.ParseBool(value)
            updateBookRead(book_title, value_as_bool)
        case "owned":
            value_as_bool, _ := strconv.ParseBool(value)
            updateBookOwned(book_title, value_as_bool)
        case "genre":
            updateBookGenre(book_title, value)
        case "entryowner":
            updateBookEntryOwner(book_title, value)
        default:
            fmt.Print("Option not set up.")
    }
}  

func updateBookTitle(book_title, title_value string) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            my_books[i].Title = title_value
            break
        }
    }
    
    writeBooks(books)
}

func updateBookSeries(book_title, series_value string) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            my_books[i].Series = series_value
            break
        }
    }
    
    writeBooks(books)
}

func updateBookAuthor(book_title, author_value string) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            my_books[i].Author = author_value
            break
        }
    }
    
    writeBooks(books)
}

func updateBookReadingList(book_title, reading_list_value string) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            my_books[i].ReadingList = reading_list_value
            break
        }
    }
    
    writeBooks(books)
}

func updateBookRecommendedBy(book_title, recommended_by_value string) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            my_books[i].RecommendedBy = recommended_by_value
            break
        }
    }
    
    writeBooks(books)
}

func updateBookRead(book_title string, read_value bool) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            my_books[i].Read = read_value
            break
        }
    }
    
    writeBooks(books)
}

func updateBookOwned(book_title string, owned_value bool) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            my_books[i].Owned = owned_value
            break
        }
    }
    
    writeBooks(books)
}

func updateBookGenre(book_title, genre_value string) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            my_books[i].Genre = genre_value
            break
        }
    }
    
    writeBooks(books)
}

func updateBookEntryOwner(book_title, entry_owner_value string) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            my_books[i].EntryOwner = entry_owner_value
            break
        }
    }
    
    writeBooks(books)
}

func deleteBook(book_title string) {
    var books Books
    books = readBooks()
    
    var my_books []Book = books.Books

    i := 0
    for ; i < len(my_books); i++ {
        if strings.ToLower(my_books[i].Title) == strings.ToLower(book_title) {
            break
        }
    }

    books.Books = append(books.Books[:i], books.Books[i+1:]...)
    
    writeBooks(books)
}
