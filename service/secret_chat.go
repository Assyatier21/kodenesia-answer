package service

import "strings"

func SecretChat(s string) string {
	var (
		words []string
		res   string
	)

	words = strings.Split(s, " ")

	for _, word := range words {
		for i := len(word) - 1; i >= 0; i-- {
			res += string(word[i])
		}
		res += " "
	}

	return res
}
