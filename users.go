package main

import (
    "encoding/json"
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

func getUser(username string) *User {
    var users []User = readUsers()

    for i := 0; i < len(users); i++ {
        if users[i].Username == username {
            return &users[i]
        }
    }

    return nil
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

    var new_user *User = &User{
        Username: username,
        ReadingLists: nil, //??? Not sure yet
    }

    appendUser(new_user)    
}

func deleteUser(username string) {
    var users []User = readUsers()

    i := 0
    for ; i < len(users); i++ {
        if strings.ToLower(users[i].Username) == strings.ToLower(username) {
            break
        }
    }

    users = append(users[:i], users[i+1:]...)
    
    writeUsers(&users)
}
