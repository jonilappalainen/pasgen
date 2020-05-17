package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    listLength := 10
    passwordLength := 10

    rand.Seed(time.Now().UnixNano())

    for i:=0; i<listLength;i++ {
        p := ""
        for j:=0; j<passwordLength;j++ {
            p += fmt.Sprintf("%x", rand.Intn(100))
        }
        fmt.Printf("%s\n", p)
    }
}
