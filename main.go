package main

import (
    "fmt"
    "math/rand"
    "time"
    "strings"
)

func GeneratePassword(passwordLength int) string {
    rand.Seed(time.Now().UnixNano())
    allowedChars := strings.Split(
        "@!#¤%€+-0123456789abcdefghijklmnopqrstuvxyzåäöABCDEFGHIJKLMNOPQRSTUVXYZÅÄÖ", "")
    p := ""
    for j:=0; j<passwordLength;j++ {
        p += allowedChars[rand.Intn(len(allowedChars))]
    }
    return p
}

func main() {
    listLength := 10
    passwordLength := 10
    for i:=0; i<listLength;i++ {
        fmt.Printf("%s\n", GeneratePassword(passwordLength))
    }
}
