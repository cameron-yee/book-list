package main

import (
    "bufio"
    "fmt"
    "strings"
    "os"
)

func colorPrintString(label, str string) {
    fmt.Printf("\033[1;35m%s\033[0m: %s\n", label, str)
}

func colorPrintBool(label string,  boolean bool) {
    fmt.Printf("\033[1;35m%s\033[0m: %t\n", label, boolean)
}

func getInput(label string) string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("%s: ", label)
    input, _ := reader.ReadString('\n')
    return strings.TrimSuffix(input, "\n")
}
