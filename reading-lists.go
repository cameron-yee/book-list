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
    data, err := ioutil.ReadFile(fmt.Sprintf("%s/reading-lists.json", getCallDirectory()))
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
            
func printReadingList(reading_list *ReadingList, verbose bool) {
    colorPrintField("Title", (*reading_list).Title)
    colorPrintField("Members", strings.Join((*reading_list).Members[:], ", "))

    var books []Book = readBooks()

    colorPrintField("Books", "")
    if len((*reading_list).Books) == 0 {
        fmt.Println("NO BOOKS IN READING LIST.") 
    } else {
        fmt.Println("-------------------------------------------------------------")
        fmt.Println("-------------------------------------------------------------")
    }
    
    for i := 0; i < len((*reading_list).Books); i++ {
        for j := 0; j < len(books); j++ {
            if strings.ToLower(books[j].Title) == strings.ToLower((*reading_list).Books[i]) {
                printBook(&books[j], true, verbose)
                break
            }
        }
    }
    
    fmt.Println("-------------------------------------------------------------")
}

func listReadingLists(verbose bool) {
    var reading_lists []ReadingList = readReadingLists()

    for i := 0; i < len(reading_lists); i++ {
        printReadingList(&reading_lists[i], verbose)
    }
}

func writeReadingLists(reading_lists *[]ReadingList) {
    gitPullOrigin()
    
    dataBytes, err := json.Marshal((*reading_lists))
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(fmt.Sprintf("%s/reading-lists.json", getCallDirectory()), dataBytes, 0644)
    if err != nil {
        panic(err)
    }
    
    gitCommit("reading-lists.json", "Edit reading-lists.json")
    gitPush()
}

func appendReadingList(readinglist *ReadingList) {
    var readinglists []ReadingList = readReadingLists()
    
    readinglists = append(readinglists, *readinglist)

    writeReadingLists(&readinglists)
}

func addReadingList() {
    title := getInput("Title")

    var reading_list_index = getReadingListIndex(title)
    if reading_list_index != -1 {
        fmt.Printf("Reading list already exists with title: \"%s\".\n", title)
        return
    } 
    
    username := getInput("Username")

    var user_index int = getUserIndex(username)
    if user_index == -1 {
        fmt.Printf("No user found with username: %s\n", username)
    }

    var users []User = readUsers()
   
    members := []string{users[user_index].Username}

    var new_reading_list *ReadingList = &ReadingList{
        Title: title,
        Members: members,
        Books: nil,
    }

    appendReadingList(new_reading_list)
}

func updateReadingListTitle(reading_list_title, new_title string) {
    var readinglists []ReadingList = readReadingLists()
    
    var updated_reading_list_index = getReadingListIndex(new_title)
    if updated_reading_list_index != -1 {
        fmt.Printf("Reading list already exists with title: \"%s\".\n", new_title)
        return
    }

    var reading_list_index = getReadingListIndex(reading_list_title)
    if reading_list_index != -1 {
        readinglists[reading_list_index].Title = new_title
        writeReadingLists(&readinglists)
    } else {
        fmt.Printf("Can't find reading list with title: \"%s\".\n", reading_list_title)
    }
}

func addBookToReadingList() {
    title := getInput("Reading List Title")
    book_title := getInput("Book Title")

    var readinglists []ReadingList = readReadingLists()

    var readinglist_index int = getReadingListIndex(title)
    if readinglist_index == -1 {
        fmt.Printf("No reading list with title: %s/n", title)
        return
    }

    var book_index int = getBookIndex(book_title)
    if book_index == -1 {
        fmt.Printf("No book found with title: %s\n", book_title)
        return
    }
    
    var book_already_in_list bool
    for i := 0; i < len(readinglists[readinglist_index].Books); i++ {
        if readinglists[readinglist_index].Books[i] == book_title {
            book_already_in_list = true
        }    
    }

    if book_already_in_list == false {
        readinglists[readinglist_index].Books = append(readinglists[readinglist_index].Books, book_title)
        writeReadingLists(&readinglists)
    } else {
        fmt.Printf("Book with title: \"%s\" already exists in reading list \"%s\".\n", book_title, title)
    }
}

func addMemberToReadingList() {
    title := getInput("Reading List Title")
    new_member := getInput("New Member Username")

    var readinglists []ReadingList = readReadingLists()
    var users []User = readUsers()

    var readinglist_index int = getReadingListIndex(title)
    if readinglist_index == -1 {
        fmt.Printf("No reading list with title: \"%s\".\n", title)
        return
    }

    var user_index int = getUserIndex(new_member)
    if user_index == -1 {
        fmt.Printf("No user found with username: \"%s\".\n", new_member)
        return
    }
    
    var user_already_in_list bool
    for i := 0; i < len(readinglists[readinglist_index].Members); i++ {
        if readinglists[readinglist_index].Members[i] == new_member {
            user_already_in_list = true
        }    
    }

    if user_already_in_list == false {
        readinglists[readinglist_index].Members = append(readinglists[readinglist_index].Members, users[user_index].Username)
        users[user_index].ReadingLists = append(users[user_index].ReadingLists, title)

        writeReadingLists(&readinglists)
        writeUsers(&users)
    } else {
        fmt.Printf("User with title: \"%s\" already exists in reading list \"%s\".\n", new_member, title)
    }
}

func deleteBookFromReadingList() {
    title := getInput("Reading List Title")
    book_title := getInput("Book Title")
    
    var readinglists []ReadingList = readReadingLists()

    var reading_list_index int = getReadingListIndex(title)
    if reading_list_index == -1 {
        fmt.Printf("No reading list with title: \"%s\".\n", title)
        return
    }

    var book_exists int = getBookIndex(book_title)
    if book_exists == -1 {
        fmt.Printf("No book found with title: \"%s\".\n", book_title)
        return
    }

    for i := 0; i < len(readinglists[reading_list_index].Books); i++ {
        if strings.ToLower(readinglists[reading_list_index].Books[i]) == strings.ToLower(book_title) {
            readinglists[reading_list_index].Books = append(readinglists[reading_list_index].Books[:i], readinglists[reading_list_index].Books[i+1:]...)
            break
        }
    }

    writeReadingLists(&readinglists)
}

func deleteMemberFromReadingList() {
    title := getInput("Reading List Title")
    member := getInput("Member Username")
    
    var readinglists []ReadingList = readReadingLists()
    var users []User = readUsers()
    
    var i int = getReadingListIndex(title)
    if i == -1 {
        fmt.Printf("No reading list with title: \"%s\".\n", title)
        return
    }

    var user_index int = getUserIndex(member)
    if user_index == -1 {
        fmt.Printf("No user found with username: \"%s\".\n", member)
        return
    }

    for j := 0; j < len(readinglists[i].Members); j++ {
        if users[user_index].Username == readinglists[i].Members[j] {
            readinglists[i].Members = append(readinglists[i].Members[:j], readinglists[i].Members[j+1:]...)
            break
        }
    }
    
    for j := 0; j < len(users[user_index].ReadingLists); j++ {
        if users[user_index].ReadingLists[j] == title {
            users[user_index].ReadingLists = append(users[user_index].ReadingLists[:j], users[user_index].ReadingLists[j+1:]...)
            break
        }
    }

    writeReadingLists(&readinglists)
    writeUsers(&users)
}

func deleteReadingList(reading_list_title string) {
    var reading_lists []ReadingList = readReadingLists()
    var users []User = readUsers()

    var reading_list_index int = getReadingListIndex(reading_list_title)
    if reading_list_index == -1 {
        fmt.Printf("No reading list with title: \"%s\".\n", reading_list_title)
        return
    }

    for i := 0; i < len(users); i++ {
        for j := 0; j < len(users[i].ReadingLists); j++ {
            if users[i].ReadingLists[j] == reading_list_title {
                users[i].ReadingLists = append(users[i].ReadingLists[:j], users[i].ReadingLists[j+1:]...)
                break
            }
        }
    }

    reading_lists = append(reading_lists[:reading_list_index], reading_lists[reading_list_index+1:]...)
    
    writeReadingLists(&reading_lists)
    writeUsers(&users)
}
