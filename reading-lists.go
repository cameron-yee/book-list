package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
)

type ReadingList struct {
    Title   string   `json:"title"`
    Members []string `json:"members"`
    Books   []string   `json:"books"`
}

func readReadingLists() []ReadingList {
    data, err := ioutil.ReadFile("./reading-lists.json")
    if err != nil {
        panic(err)
    }

    var readinglists []ReadingList

    err = json.Unmarshal(data, &readinglists)
    if err != nil {
        panic(err)
    }

    return readinglists
}

func getReadingListIndex(reading_list_title string) int {
    var reading_lists []ReadingList = readReadingLists()

    for i := 0; i < len(reading_lists); i++ {
        if strings.ToLower(reading_lists[i].Title) == strings.ToLower(reading_list_title) {
            return i
        }
    }

    return -1
}

func runUpdateReadingList() {
    reading_list_title := getInput("Reading List Title")
    field := getInput("Field")
    value := getInput("New Value")
    
    switch strings.ToLower(field) {
        case "title":
            updateReadingListTitle(reading_list_title, value)
        default:
            fmt.Println("Options are title.")
    }
}
            
func printReadingList(reading_list ReadingList) {
    colorPrintString("Title", reading_list.Title)
    colorPrintString("Members", strings.Join(reading_list.Members[:], ", "))

    var books []Book = readBooks()

    colorPrintString("Books", "")
    fmt.Println("-------------------------------------------------------------")
    for i := 0; i < len(reading_list.Books); i++ {
        for j := 0; j < len(books); j++ {
            if strings.ToLower(books[j].Title) == strings.ToLower(reading_list.Books[i]) {
                printBook(books[j])
                break
            }
        }
    }
    
    fmt.Println("*************************************************************")
    fmt.Println("*************************************************************")
    fmt.Println("*************************************************************")
}

func listReadingLists() {
    var reading_lists []ReadingList = readReadingLists()

    for i := 0; i < len(reading_lists); i++ {
        printReadingList(reading_lists[i])
    }
}

func writeReadingLists(reading_lists *[]ReadingList) {
    dataBytes, err := json.Marshal((*reading_lists))
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile("./reading-lists.json", dataBytes, 0644)
    if err != nil {
        panic(err)
    }
}

func appendReadingList(readinglist *ReadingList) {
    var readinglists []ReadingList = readReadingLists()
    
    readinglists = append(readinglists, *readinglist)

    writeReadingLists(&readinglists)
}

func addReadingList() {
    title := getInput("Title")
    username := getInput("Username")

    var user *User = getUser(username)
    if user == nil {
        fmt.Printf("No user found with username: %s\n", username))
    }
   
    members := []string{(*user).Username}

    var new_reading_list *ReadingList = &ReadingList{
        Title: title,
        Members: members,
        Books: nil,
    }

    appendReadingList(new_reading_list)
}

func updateReadingListTitle(reading_list_title, new_title string) {
    var readinglists []ReadingList = readReadingLists()

    for i := 0; i < len(readinglists); i++ {
        if strings.ToLower(readinglists[i].Title) == strings.ToLower(reading_list_title) {
           readinglists[i].Title = new_title
           break
        }
    }

    writeReadingLists(&readinglists)
}

func addBookToReadingList() {
    title := getInput("Reading List Title")
    book_title := getInput("Book Title")

    var readinglists []ReadingList = readReadingLists()

    var i int = getReadingListIndex(title)
    if i == -1 {
        fmt.Printf("No reading list with title: %s/n", title)
        return
    }

    var book_index int = getBookIndex(book_title)
    if book_index == -1 {
        fmt.Printf("No book found with title: %s\n", book_title)
        return
    }

    var books []Book = readBooks()
   
    readinglists[i].Books = append(readinglists[i].Books, books[book_index].Title)

    writeReadingLists(&readinglists)
}

func addMemberToReadingList() {
    title := getInput("Reading List Title")
    new_member := getInput("New Member Username")

    var readinglists []ReadingList = readReadingLists()

    var i int = getReadingListIndex(title)
    if i == -1 {
        fmt.Printf("No reading list with title: %s/n", title)
        return
    }

    var user *User = getUser(new_member)
    if user == nil {
        fmt.Printf("No user found with username: %s", new_member)
        return
    }
   
    readinglists[i].Members = append(readinglists[i].Members, (*user).Username)

    writeReadingLists(&readinglists)
}

func deleteBookFromReadingList() {
    title := getInput("Reading List Title")
    book_title := getInput("Book Title")
    
    var readinglists []ReadingList = readReadingLists()

    var i int = getReadingListIndex(title)
    if i == -1 {
        fmt.Printf("No reading list with title: %s/n", title)
        return
    }

    var book_exists int = getBookIndex(book_title)
    if book_exists == -1 {
        fmt.Printf("No book found with title: %s", book_title)
        return
    }

    for j := 0; j < len(readinglists[i].Books); j++ {
        if strings.ToLower(readinglists[i].Books[j]) == strings.ToLower(book_title) {
            readinglists[i].Books = append(readinglists[i].Books[:j], readinglists[i].Books[j+1:]...)
            break
        }
    }

    writeReadingLists(&readinglists)
}

func deleteMemberFromReadingList() {
    title := getInput("Reading List Title")
    member := getInput("New Member Username")
    
    var readinglists []ReadingList = readReadingLists()
    
    var i int = getReadingListIndex(title)
    if i == -1 {
        fmt.Printf("No reading list with title: %s/n", title)
        return
    }

    var user *User = getUser(member)
    if user == nil {
        fmt.Printf("No user found with username: %s", member)
        return
    }

    for j := 0; j < len(readinglists[i].Members); j++ {
        if (*user).Username == readinglists[i].Members[j] {
            readinglists[i].Members = append(readinglists[i].Members[:j], readinglists[i].Members[j+1:]...)
            break
        }
    }

    writeReadingLists(&readinglists)
}

func deleteReadingList(reading_list_title string) {
    var reading_lists []ReadingList = readReadingLists()

    var i int = getReadingListIndex(reading_list_title)

    reading_lists = append(reading_lists[:i], reading_lists[i+1:]...)
    
    writeReadingLists(&reading_lists)
}
