package pasgen

import (
    "math/rand"
    "time"
    "strings"
)

func GeneratePassword(passwordLength int) string {
    rand.Seed(time.Now().UnixNano())
    allowedChars := strings.Split(
        "@!#%+-0123456789abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ", "")
    p := ""
    for j:=0; j<passwordLength;j++ {
        p += allowedChars[rand.Intn(len(allowedChars))]
    }
    return p
}
