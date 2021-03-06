package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

type Book struct {
    Title         string    `json:"title"`
    Series        string    `json:"series"`
    Author        string    `json:"author"`
    RecommendedBy string    `json:"recommendedBy"`
    ReadBy        []string  `json:"readBy"`
    Owned         bool      `json:"owned"`
    Genre         string    `json:"genre"`
    EntryOwner    string    `json:"entryOwner"`
}

func readBooks() []Book {
    data, err := ioutil.ReadFile(fmt.Sprintf("%s/books.json", getCallDirectory()))
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
        if compareStringsCaseInsensitive(books[i].Title, book_title) {
            return i
        }
    }

    return -1
}

func printCantFindBook(book_title string) {
    fmt.Printf("Can't find book with title \"%s\".\n", book_title)
}
 
func printBook(book *Book, indent bool, verbose bool, vverbose bool) {
    var prefix string

    if indent {
        prefix = "\t"
    }
    
    colorPrintField(fmt.Sprintf("%sTitle", prefix), (*book).Title)

    if verbose || vverbose {
        colorPrintField(fmt.Sprintf("%sSeries", prefix), (*book).Series)
        colorPrintField(fmt.Sprintf("%sAuthor", prefix), (*book).Author)
        colorPrintField(fmt.Sprintf("%sGenre", prefix), (*book).Genre)
    }

    if vverbose {
        colorPrintField(fmt.Sprintf("%sRecommended By", prefix), (*book).RecommendedBy)
        colorPrintField(fmt.Sprintf("%sReadBy", prefix), strings.Join((*book).ReadBy[:], ", "))
        colorPrintField(fmt.Sprintf("%sOwned", prefix), strconv.FormatBool((*book).Owned))
        colorPrintField(fmt.Sprintf("%sEntryOwner", prefix), (*book).EntryOwner)
    }
    
    fmt.Println("-------------------------------------------------------------")
}

func listBooks(verbose bool, vverbose bool, limit int) {
    var books []Book = readBooks()

    var until int = len(books)
    if limit != 0 {
        until = limit
    }

    for i := 0; i < until; i++ {
        printBook(&books[i], false, verbose, vverbose)
    }
}

func writeBooks(books *[]Book) {
    gitPullOrigin(true)
    
    dataBytes, err := json.Marshal((*books))
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(fmt.Sprintf("%s/books.json", getCallDirectory()), dataBytes, 0644)
    if err != nil {
        panic(err)
    }

    gitCommit("books.json", "Edit books.json")
    gitPush()
}

func appendBook(book *Book) {
    var books []Book = readBooks()
    
    books = append(books, *book)

    writeBooks(&books)
}

func addBookReadBy(read_by *[]string) {
    var cont bool = true
    for cont {
        person := getInput("Person: ")
        (*read_by) = append((*read_by), person)
        
        cont_prompt := getInput("Add another person? (y/n) ")
        var cont_strings []string = []string{"y", "yes", ""}

        cont = false
        for i := 0; i < len(cont_strings); i++ {
            if compareStringsCaseInsensitive(cont_prompt, cont_strings[i]) {
                cont = true
            } 
        }
    }
}

func deleteBookReadBy(read_by *[]string) {
    var cont bool = true
    for cont {
        person := getInput("Person: ")

        for i := 0; i < len((*read_by)); i++ {
            if compareStringsCaseInsensitive(person, (*read_by)[i]) {
                (*read_by) = append((*read_by)[:i], (*read_by)[i+1:]...)
            }
        }

        if len((*read_by)) == 0 {
            return
        }
        
        cont_prompt := getInput("Delete another person? (y/n) ")
        var cont_strings []string = []string{"y", "yes", ""}

        cont = false
        for i := 0; i < len(cont_strings); i++ {
            if compareStringsCaseInsensitive(cont_prompt, cont_strings[i]) {
                cont = true
            } 
        }
    }
}

func addBook() {
    title := getInput("Title: ")

    var book_index int = getBookIndex(title)
    
    if book_index != -1 {
        fmt.Printf("Book already exists with title: \"%s\".\n", title)
        return
    }
    
    series := getInput("Series: ")
    author := getInput("Author: ")
    recommended_by := getInput("Recommended By: ")

    fmt.Println("Add people that have read this book:")
    var read_by []string
    addBookReadBy(&read_by)
    
    owned := getInput("Owned (t/f): ")
    owned_bool := false
    if strings.ToLower(owned) == "true" || strings.ToLower(owned) == "t" {
        owned_bool = true    
    }
    
    genre := getInput("Genre: ")
    entry_owner := getInput("Entry Owner: ")

    if entry_owner == "" {
        GITHUB_USER, _ := os.LookupEnv("GITHUB_USER")
        entry_owner = GITHUB_USER

        var users []User = readUsers()
        
        for i := 0; i < len(users); i++ {
            if users[i].GitHubUser == GITHUB_USER {
                entry_owner = users[i].Username
                break
            }
        }
    }
    
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
        ReadBy: read_by,
        Owned: owned_bool,
        Genre: genre,
        EntryOwner: entry_owner,
    }

    appendBook(new_book)
}

func updateBookTitle(book_title, new_book_title string) {
    var new_book_index = getBookIndex(new_book_title)
    
    if new_book_index != -1 {
        fmt.Printf("Book already exists with title: \"%s\".\n", new_book_title)
        return
    }
    
    var book_index int = getBookIndex(book_title)
    if book_index == -1 {
        printCantFindBook(book_title)
        return
    }
    
    var books []Book = readBooks()
    books[book_index].Title = new_book_title
    writeBooks(&books)
    
    var reading_lists []ReadingList = readReadingLists()

    for i := 0; i < len(reading_lists); i++ {
        for j := 0; j < len(reading_lists[i].Books); j++ {
            if compareStringsCaseInsensitive(reading_lists[i].Books[j], book_title) {
                reading_lists[i].Books[j] = new_book_title
            }
        }
    }

    writeReadingLists(&reading_lists)
}

func updateBookSeries(book_title, series_value string) {
    var book_index int = getBookIndex(book_title)

    if book_index == -1 {
        printCantFindBook(book_title)
        return
    }
    
    var books []Book = readBooks()
    books[book_index].Series = series_value
    writeBooks(&books)

}

func updateBookAuthor(book_title, author_value string) {
    var books []Book = readBooks()
    
    var book_index int = getBookIndex(book_title)
    if book_index == -1 {
        printCantFindBook(book_title)
        return
    }

    books[book_index].Author = author_value
   
    writeBooks(&books)
}

func updateBookRecommendedBy(book_title, recommended_by_value string) {
    var book_index int = getBookIndex(book_title)
    
    if book_index == -1 {
        printCantFindBook(book_title)
        return
    }
    
    var books []Book = readBooks()
    books[book_index].RecommendedBy = recommended_by_value
    writeBooks(&books)
}

func updateBookOwned(book_title string, owned_value bool) {
    var book_index int = getBookIndex(book_title)
    
    if book_index == -1 {
        printCantFindBook(book_title)
        return
    }
    
    var books []Book = readBooks()
    books[book_index].Owned = owned_value
    writeBooks(&books)
}

func updateBookGenre(book_title, genre_value string) {
    var book_index int = getBookIndex(book_title)
    
    if book_index == -1 {
        printCantFindBook(book_title)
        return
    }
    
    var books []Book = readBooks()
    books[book_index].Genre = genre_value
    writeBooks(&books)
}

func updateBookEntryOwner(book_title, entry_owner_value string) {
    var book_index int = getBookIndex(book_title)
    
    if book_index == -1 {
        printCantFindBook(book_title)
        return
    }
    
    var books []Book = readBooks()
    books[book_index].EntryOwner = entry_owner_value
    writeBooks(&books)
}

func runUpdateBook() {
    book_title := getInput("Book Title: ")
    field := getInput("Field: ")
    
    switch strings.ToLower(field) {
        case "title":
            value := getInput("New Value: ")
            updateBookTitle(book_title, value)
        case "series":
            value := getInput("New Value: ")
            updateBookSeries(book_title, value)
        case "author":
            value := getInput("New Value: ")
            updateBookAuthor(book_title, value)
        case "recommendedby":
            value := getInput("New Value: ")
            updateBookRecommendedBy(book_title, value)
        case "readby":
            add_or_delete := getInput("Add or Delete? (a/d) ")
            var books []Book = readBooks()
            var i int = getBookIndex(book_title)

            if strings.ToLower(add_or_delete) == "Add" || strings.ToLower(add_or_delete) == "a" {
                addBookReadBy(&books[i].ReadBy)               
                writeBooks(&books)
            } else if strings.ToLower(add_or_delete) == "Delete" || strings.ToLower(add_or_delete) == "d" {
                deleteBookReadBy(&books[i].ReadBy)               
                writeBooks(&books)
            } else {
                fmt.Printf("Select add, delete, a, or d.")
            }
        case "owned":
            value := getInput("New Value: ")
            value_as_bool, _ := strconv.ParseBool(value)
            updateBookOwned(book_title, value_as_bool)
        case "genre":
            value := getInput("New Value: ")
            updateBookGenre(book_title, value)
        case "entryowner":
            value := getInput("New Value: ")
            updateBookEntryOwner(book_title, value)
        default:
            fmt.Println("Options are title, series, author, recommendedby, readby, owned, genre, or entryowner.")
    }
}  

func deleteBook(book_title string) {
    var book_index int = getBookIndex(book_title)
    
    if book_index == -1 {
        printCantFindBook(book_title)
        return
    }
    
    var books []Book = readBooks()
    books = append(books[:book_index], books[book_index+1:]...)
    
    writeBooks(&books)
}
