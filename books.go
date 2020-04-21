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

func printCantFindBook(book_title string) {
    fmt.Printf("Can't find book with title \"%s\".\n", book_title)
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

    var book_index int = getBookIndex(title)
    
    if book_index != -1 {
        fmt.Printf("Book with title \"%s\" already exists.", title)
        return
    }
    
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

func updateBookTitle(book_title, new_book_title string) {
    var new_book_index = getBookIndex(new_book_title)
    
    if new_book_index != -1 {
        fmt.Printf("Book with title \"%s\" already exists.", new_book_title)
        return
    }
    
    var book_index int = getBookIndex(book_title)
    if book_index != -1 {
        var books []Book = readBooks()
        books[book_index].Title = new_book_title
        writeBooks(&books)
    } else {
        printCantFindBook(book_title)
        return
    }
    
    var reading_lists []ReadingList = readReadingLists()

    for i := 0; i < len(reading_lists); i++ {
        for j := 0; j < len(reading_lists[i].Books); j++ {
            if strings.ToLower(reading_lists[i].Books[j]) == strings.ToLower(book_title) {
                reading_lists[i].Books[j] = new_book_title
            }
        }
    }

    writeReadingLists(&reading_lists)
}

func updateBookSeries(book_title, series_value string) {
    var book_index int = getBookIndex(book_title)
    
    if book_index != -1 {
        var books []Book = readBooks()
        books[book_index].Series = series_value
        writeBooks(&books)
    } else {
        printCantFindBook(book_title)
    }
}

func updateBookAuthor(book_title, author_value string) {
    var books []Book = readBooks()
    
    var book_index int = getBookIndex(book_title)
    if book_index != -1 {
        books[book_index].Author = author_value
    } else {
        printCantFindBook(book_title)
    }
    
    writeBooks(&books)
}

func updateBookRecommendedBy(book_title, recommended_by_value string) {
    var book_index int = getBookIndex(book_title)
    
    if book_index != -1 {
        var books []Book = readBooks()
        books[book_index].RecommendedBy = recommended_by_value
        writeBooks(&books)
    } else {
        printCantFindBook(book_title)
    }

}

func updateBookRead(book_title string, read_value bool) {
    var book_index int = getBookIndex(book_title)
    
    if book_index != -1 {
        var books []Book = readBooks()
        books[book_index].Read = read_value
        writeBooks(&books)
    } else {
        printCantFindBook(book_title)
    }
}

func updateBookOwned(book_title string, owned_value bool) {
    var book_index int = getBookIndex(book_title)
    
    if book_index != -1 {
        var books []Book = readBooks()
        books[book_index].Owned = owned_value
        writeBooks(&books)
    } else {
        printCantFindBook(book_title)
    }
}

func updateBookGenre(book_title, genre_value string) {
    var book_index int = getBookIndex(book_title)
    
    if book_index != -1 {
        var books []Book = readBooks()
        books[book_index].Genre = genre_value
        writeBooks(&books)
    } else {
        printCantFindBook(book_title)
    }
}

func updateBookEntryOwner(book_title, entry_owner_value string) {
    var book_index int = getBookIndex(book_title)
    
    if book_index != -1 {
        var books []Book = readBooks()
        books[book_index].EntryOwner = entry_owner_value
        writeBooks(&books)
    } else {
        printCantFindBook(book_title)
    }
}

func deleteBook(book_title string) {
    var book_index int = getBookIndex(book_title)
    
    if book_index != -1 {
        var books []Book = readBooks()
        books = append(books[:book_index], books[book_index+1:]...)
        writeBooks(&books)
    } else {
        printCantFindBook(book_title)
    }
}
