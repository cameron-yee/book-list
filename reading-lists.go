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
    Books   []Book   `json:"books"`
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

func printReadingList(reading_list ReadingList) {
    colorPrintString("Title", reading_list.Title)
    colorPrintString("Members", strings.Join(reading_list.Members[:], ", "))

    for i := 0; i < len(reading_list.Books); i++ {
        printBook(reading_list.Books[i])
    }
    
    fmt.Println("-------------------------------------------------------------")
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
        panic(fmt.Sprintf("No user found with username: %s", username))
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

func deleteReadingList(reading_list_title string) {
    var reading_lists []ReadingList = readReadingLists()

    i := 0
    for ; i < len(reading_lists); i++ {
        if strings.ToLower(reading_lists[i].Title) == strings.ToLower(reading_list_title) {
            break
        }
    }

    reading_lists = append(reading_lists[:i], reading_lists[i+1:]...)
    
    writeReadingLists(&reading_lists)
}
