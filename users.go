package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
)

type User struct {
    GitHubUser     string   `json:"githubuser"`
    Username       string   `json:"username"`
    ReadingLists   []string `json:"readingLists"` 
}

func readUsers() []User {
    data, err := ioutil.ReadFile(fmt.Sprintf("%s/users.json", getCallDirectory()))
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

func printUser(user *User, verbose bool) {
    colorPrintField("Username", (*user).Username)

    if verbose {
        colorPrintField("GitHubUser", (*user).GitHubUser)
        colorPrintField("Reading Lists", strings.Join((*user).ReadingLists[:], ", "))
    }
    
    fmt.Println("-------------------------------------------------------------")
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

func listUsers(verbose bool, limit int) {
    var users []User = readUsers()
    
    var until = len(users)
    if limit != 0 {
        until = limit
    }


    for i := 0; i < until; i++ {
        if verbose {
            printUser(&users[i], true)
        } else {
            printUser(&users[i], false)
        }
    }
}


func writeUsers(users *[]User) {
    gitPullOrigin(true)
    
    dataBytes, err := json.Marshal((*users))
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(fmt.Sprintf("%s/users.json", getCallDirectory()), dataBytes, 0644)
    if err != nil {
        panic(err)
    }
    
    gitCommit("users.json", "Edit users.json")
    gitPush()
}

func appendUser(user *User) {
    var users []User = readUsers()
    
    users = append(users, *user)

    writeUsers(&users)
}

func addUser() {
    username := getInput("Username: ")
    githubuser := getInput("GitHubUser: ")

    var user_index int = getUserIndex(username)
    if user_index != -1 {
        fmt.Printf("User already exists with username: \"%s\".\n", username)
        return
    }

    var new_user *User = &User{
        GitHubUser: githubuser,
        Username: username,
        ReadingLists: nil, //??? Not sure yet
    }

    appendUser(new_user)
}

func updateUserGitHubUser(username, new_githubuser string) {
    var user_index int = getUserIndex(username)
    if user_index == -1 {
        fmt.Printf("User doesn't exist with username: \"%s\".\n", username)
        return        
    }
    
    var users []User = readUsers()
    users[user_index].GitHubUser = new_githubuser
    writeUsers(&users)
}

func updateUserUsername(username, new_username string) {
    var new_username_index int = getUserIndex(new_username)
    if new_username_index != -1 {
        fmt.Printf("User already exists with username: \"%s\".\n", new_username)
        return
    }
    
    var user_index int = getUserIndex(username)
    if user_index == -1 {
        fmt.Printf("User doesn't exist with username: \"%s\".\n", username)
        return
    }
    
    var users []User = readUsers()
    users[user_index].Username = new_username
    writeUsers(&users)

    var books []Book = readBooks()

    for i := 0; i < len(books); i++ {
        if books[i].EntryOwner == username {
            books[i].EntryOwner = new_username
        } 
    }

    writeBooks(&books)
}

func runUpdateUser() {
    username := getInput("Username: ")
    field := getInput("Field: ")
    value := getInput("New Value: ")
    
    switch strings.ToLower(field) {
        case "username":
            updateUserUsername(username, value)
        case "githubuser":
            updateUserGitHubUser(username, value)
        default:
            fmt.Println("Options are username and githubuser.")
    }
}

func deleteUser(username string) {
    var user_index int = getUserIndex(username)
    
    if user_index == -1 {
        fmt.Printf("User doesn't exist with username: \"%s\".\n", username)
        return
    }
    
    var users []User = readUsers()
    users = append(users[:user_index], users[user_index+1:]...)
    writeUsers(&users)
}
