package internal

import (
	"math/rand"
)

func GeneratePassword(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var password = []rune{}

	for range length {
		password = append(password, letters[rand.Intn(len(letters))])
	}

	return string(password)
}
