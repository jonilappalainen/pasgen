package main

import (
    "testing"
    "unicode/utf8"
)

func verifyLength(t *testing.T, password string, expected int) {
    actual := utf8.RuneCountInString(password)
    if (actual != expected) {
        t.Errorf("unexpected length: %d (expected %d) '%s'", actual, expected, password)
    }
}

func TestGeneratedPasswordLength(t *testing.T) {
    verifyLength(t, GeneratePassword(3), 3)
    verifyLength(t, GeneratePassword(0), 0)
    verifyLength(t, GeneratePassword(1), 1)
    verifyLength(t, GeneratePassword(10), 10)
    verifyLength(t, GeneratePassword(100), 100)
}
