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
