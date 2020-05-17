package main

import (
    "os"
    "fmt"
    "math/rand"
    "time"
    "strings"
    "strconv"
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

func GetEnv(key string) string {
  return os.Getenv(key)
}

func GetEnvAsInteger(key string) int {
    i, _ := strconv.Atoi(GetEnv(key))
    return i
}

func main() {
    listLength := GetEnvAsInteger("PASSWORD_COUNT")
    passwordLength := GetEnvAsInteger("PASSWORD_LENGTH")

    for i:=0; i<listLength;i++ {
        fmt.Printf("%s\n", GeneratePassword(passwordLength))
    }
}
