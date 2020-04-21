package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
)

type User struct {
    Username       string `json:"username"`
    ReadingLists   []string `json:"readingLists"` 
}

func readUsers() []User {
    data, err := ioutil.ReadFile("./users.json")
    if err != nil {
        panic(err)
    }

    var users []User

    err = json.Unmarshal(data, &users)
    if err != nil {
        panic(err)
    }

    return users
}

func runUpdateUser() {
    username := getInput("Username")
    field := getInput("Field")
    value := getInput("New value")
    
    switch strings.ToLower(field) {
        case "username":
            updateUserUsername(username, value)
        default:
            fmt.Println("Options are username.")
    }
}

func printUser(user User) {
    colorPrintString("Username", user.Username)
    colorPrintString("Reading Lists", strings.Join(user.ReadingLists[:], ", "))
    
    fmt.Println("-------------------------------------------------------------")
}

func listUsers() {
    var users []User = readUsers()

    for i := 0; i < len(users); i++ {
        printUser(users[i])
    }
}

func getUserIndex(username string) int {
    var users []User = readUsers()

    for i := 0; i < len(users); i++ {
        if users[i].Username == username {
            return i
        }
    }

    return -1
}

func writeUsers(users *[]User) {
    dataBytes, err := json.Marshal((*users))
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile("./users.json", dataBytes, 0644)
    if err != nil {
        panic(err)
    }
}

func appendUser(user *User) {
    var users []User = readUsers()
    
    users = append(users, *user)

    writeUsers(&users)
}

func addUser() {
    username := getInput("Username")

    var user_index int = getUserIndex(username)
    if user_index != -1 {
        fmt.Printf("User already exists with username: \"%s\".\n", username)
        return
    }

    var new_user *User = &User{
        Username: username,
        ReadingLists: nil, //??? Not sure yet
    }

    appendUser(new_user)    
}

func updateUserUsername(username, new_username string) {
    var new_username_index int = getUserIndex(new_username)
    if new_username_index != -1 {
        fmt.Printf("User already exists with username: \"%s\".\n", new_username)
        return
    }
    
    var user_index int = getUserIndex(username)
    if user_index != -1 {
        var users []User = readUsers()
        users[user_index].Username = new_username
        writeUsers(&users)
    } else {
        fmt.Printf("User doesn't exist with username: \"%s\".\n", username)
        return
    }
    
    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if books[i].EntryOwner == username {
            books[i].EntryOwner = new_username
        } 
    }

    writeBooks(&books)
}

func deleteUser(username string) {
    var user_index int = getUserIndex(username)
    
    if user_index != -1 {
        var users []User = readUsers()
        users = append(users[:user_index], users[user_index+1:]...)
        writeUsers(&users)
    } else {
        fmt.Printf("User doesn't exist with username: \"%s\".\n", username)
    }
}
