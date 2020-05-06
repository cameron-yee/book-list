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

func compareStringsCaseInsensitive(a, b string) bool {
    return strings.ToLower(a) == strings.ToLower(b)
}

func containsCaseInsensitive(str, substr string) bool {
    return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

func removeReturnChars(s string) string {
    // Windows uses \r\n
    // MacOS uses \n
    s = strings.TrimSuffix(s, "\n")
    s = strings.TrimSuffix(s, "\r")

    return s
}

func getInput(label string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("%s", label)
    input, _ := reader.ReadString('\n')
    return removeReturnChars(input)
}

func getCallDirectory() string {
    _, filename, _, ok := runtime.Caller(0)
    if !ok {
        panic("No caller information")
    }

    var pathname string = path.Dir(filename)

    return pathname
}
