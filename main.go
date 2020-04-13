package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
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

func colorPrintString(label, str string) {
    fmt.Printf("\033[1;35m%s\033[0m: %s\n", label, str)
}

func colorPrintBool(label string,  boolean bool) {
    fmt.Printf("\033[1;35m%s\033[0m: %t\n", label, boolean)
}

func listBooks() {
    var books Books

    books = readBooks()
    
    var my_books []Book = books.Books

    for i := 0; i < len(my_books); i++ {
        colorPrintString("Title", my_books[i].Title)
        colorPrintString("Series", my_books[i].Series)
        colorPrintString("Author", my_books[i].Author)
        colorPrintString("Recommended By", my_books[i].RecommendedBy)
        colorPrintBool("Read", my_books[i].Read)
        colorPrintBool("Owned", my_books[i].Owned)
        colorPrintString("Genre", my_books[i].Genre)
        fmt.Println("-------------------------------------------------------------")
    }
}

func getInput(label string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("%s: ", label)
    input, _ := reader.ReadString('\n')
    return strings.TrimSuffix(input, "\n")
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

func help() {
    fmt.Println("Commands:")
    fmt.Println("\thelp: Print list of commands.")
    fmt.Println("\tlist: List books.")
    fmt.Println("\tadd: Add a book.")
}

func main() {
    if len(os.Args) != 2 {
        help()
        return
    }

    switch command := os.Args[1]; command {
        case "list":
            listBooks()
        case "add":
            addBook()
        default:
            help()
    }
}
