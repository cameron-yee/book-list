package main

import (
    "fmt"
    "os"
    "strings"
    "time"
    
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/plumbing/object"
    "github.com/go-git/go-git/v5/plumbing/transport/http"
)

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
    if len(os.Args) < len(arg)+1 {
        Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
        os.Exit(1)
    }
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
  if err == nil {
      return
  } else if fmt.Sprintf("%s", err) != "already up-to-date" {
      fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
      os.Exit(1)
  }
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
    fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
  fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func gitPullOrigin() {
    // We instantiate a new repository targeting the given path (the .git folder)
    r, err := git.PlainOpen(getCallDirectory())
    CheckIfError(err)
    
    //  // Get the working directory for the repository
    w, err := r.Worktree()
    CheckIfError(err)
    
    //  // Pull the latest changes from the origin remote and merge into the current branch
    Info("git pull origin")
    err = w.Pull(&git.PullOptions{RemoteName: "origin"})
    CheckIfError(err)
    
    //  // Print the latest commit that was just pulled
    //ref, err := r.Head()
    //CheckIfError(err)
    //commit, err := r.CommitObject(ref.Hash())
    //CheckIfError(err)
    
    //fmt.Println(commit)
}

func gitCommit(edited_file, commit_message string) {
    // Opens an already existing repository.
    r, err := git.PlainOpen(getCallDirectory()) //directory .git lives in
    CheckIfError(err)

    w, err := r.Worktree()
    CheckIfError(err)

    // Adds the new file to the staging area.
    Info(fmt.Sprintf("git add %s", edited_file))
    _, err = w.Add(edited_file)
    CheckIfError(err)
    // We can verify the current status of the worktree using the method Status.
    // Info("git status --porcelain")
    // status, err := w.Status()
    // CheckIfError(err)
    // fmt.Println(status)
    
    // Commits the current staging area to the repository, with the new file
    // just created. We should provide the object.Signature of Author of the
    // commit.
    NAME, _     := os.LookupEnv("NAME")
    EMAIL, _ := os.LookupEnv("EMAIL")
    
    Info(fmt.Sprintf("git commit -m \"%s\"", commit_message))
    commit, err := w.Commit(fmt.Sprintf("%s", commit_message), &git.CommitOptions{
        Author: &object.Signature{
            Name:  NAME,
            Email: EMAIL,
            When:  time.Now(),
        },
    })
    CheckIfError(err)
    
    // Prints the current HEAD to verify that all worked well.
    //Info("git show -s")
    obj, err := r.CommitObject(commit) //_ == obj
    CheckIfError(err)
    fmt.Println(obj)
}

func gitPush() {
    r, err := git.PlainOpen(getCallDirectory())
    CheckIfError(err)

    Info("git push")
    // push using default options

    GITHUB_USER, _     := os.LookupEnv("GITHUB_USER")
    GITHUB_PASSWORD, _ := os.LookupEnv("GITHUB_PASSWORD")
    
    err = r.Push(&git.PushOptions{
        Auth: &http.BasicAuth{
             Username: GITHUB_USER,
             Password: GITHUB_PASSWORD,
        },
    })
    CheckIfError(err)
}
