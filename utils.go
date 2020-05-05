package main

import (
    "bufio"
    "fmt"
    "path"
    "runtime"
    "strings"
    "os"
)

func colorPrintField(label, str string) {
    fmt.Printf("\033[1;35m%s\033[0m: %s\n", label, str)
}

func getInput(label string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("%s", label)
    input, _ := reader.ReadString('\n')
    return strings.TrimSuffix(input, "\n")
}

func getCallDirectory() string {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("No caller information")
    }

    var pathname string = path.Dir(filename)

    return pathname
}
