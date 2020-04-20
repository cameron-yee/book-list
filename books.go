package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

type Book struct {
    Title         string               `json:"title"`
    Series        string               `json:"series"`
    Author        string               `json:"author"`
    RecommendedBy string               `json:"recommendedBy"`
    Read          bool                 `json:"read"`
    Owned         bool                 `json:"owned"`
    Genre         string               `json:"genre"`
    EntryOwner    string               `json:"entryOwner"`
}

func readBooks() []Book {
    data, err := ioutil.ReadFile("./books.json")
    if err != nil {
        panic(err)
    }

    var books []Book

    err = json.Unmarshal(data, &books)
    if err != nil {
        panic(err)
    }

    return books
}

func getBookIndex(book_title string) int {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            return i
        }
    }

    return -1
}
 
func printBook(book Book) {
    colorPrintString("Title", book.Title)
    colorPrintString("Series", book.Series)
    colorPrintString("Author", book.Author)
    colorPrintString("Recommended By", book.RecommendedBy)
    colorPrintBool("Read", book.Read)
    colorPrintBool("Owned", book.Owned)
    colorPrintString("Genre", book.Genre)
    colorPrintString("EntryOwner", book.EntryOwner)
    fmt.Println("-------------------------------------------------------------")
}

func listBooks() {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        printBook(books[i])
    }
}

func writeBooks(books *[]Book) {
    dataBytes, err := json.Marshal((*books))
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile("./books.json", dataBytes, 0644)
    if err != nil {
        panic(err)
    }
}

func appendBook(book *Book) {
    var books []Book = readBooks()
    
    books = append(books, *book)

    writeBooks(&books)
}

func addBook() {
    title := getInput("Title")
    series := getInput("Series")
    author := getInput("Author")
    // reading_list := getInput("Reading List")
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
    var user_index  = getUserIndex(entry_owner)
    if user_index == -1 {
        fmt.Printf("No user found with username: %s\n", entry_owner)
        return
    }
   

    var new_book *Book = &Book{
        Title: title,
        Series: series,
        Author: author,
        RecommendedBy: recommended_by,
        Read: read_bool,
        Owned: owned_bool,
        Genre: genre,
        EntryOwner: entry_owner,
    }

    appendBook(new_book)
}

func runUpdateBook() {
    book_title := getInput("Book Title")
    field := getInput("Field")
    value := getInput("New Value")
    
    switch strings.ToLower(field) {
        case "title":
            updateBookTitle(book_title, value)
        case "series":
            updateBookSeries(book_title, value)
        case "author":
            updateBookAuthor(book_title, value)
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
            fmt.Println("Options are title, series, author, recommendedby, read, owned, genre, or entryowner.")
    }
}  

func updateBookTitle(book_title, title_value string) {
    var books []Book = readBooks()
    var reading_lists []ReadingList = readReadingLists()

    for i := 0; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            books[i].Title = title_value
            break
        }
    }

    for i := 0; i < len(reading_lists); i++ {
        for j := 0; j < len(reading_lists[i].Books); j++ {
            if strings.ToLower(reading_lists[i].Books[j]) == strings.ToLower(book_title) {
                reading_lists[i].Books[j] = title_value
            }
        }
    }

    writeBooks(&books)
    writeReadingLists(&reading_lists)
}

func updateBookSeries(book_title, series_value string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            books[i].Series = series_value
            break
        }
    }
    
    writeBooks(&books)
}

func updateBookAuthor(book_title, author_value string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            books[i].Author = author_value
            break
        }
    }
    
    writeBooks(&books)
}

func updateBookRecommendedBy(book_title, recommended_by_value string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            books[i].RecommendedBy = recommended_by_value
            break
        }
    }
    
    writeBooks(&books)
}

func updateBookRead(book_title string, read_value bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            books[i].Read = read_value
            break
        }
    }
    
    writeBooks(&books)
}

func updateBookOwned(book_title string, owned_value bool) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            books[i].Owned = owned_value
            break
        }
    }
    
    writeBooks(&books)
}

func updateBookGenre(book_title, genre_value string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            books[i].Genre = genre_value
            break
        }
    }
    
    writeBooks(&books)
}

func updateBookEntryOwner(book_title, entry_owner_value string) {
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            books[i].EntryOwner = entry_owner_value
            break
        }
    }
    
    writeBooks(&books)
}

func deleteBook(book_title string) {
    var books []Book = readBooks()

    i := 0
    for ; i < len(books); i++ {
        if strings.ToLower(books[i].Title) == strings.ToLower(book_title) {
            break
        }
    }

    books = append(books[:i], books[i+1:]...)
    
    writeBooks(&books)
}
