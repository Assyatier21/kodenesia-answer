package service

import "strings"

func RichWord(s string) int {
	var (
		words []string
		res   = 0
	)

	words = strings.Split(s, " ")

	for _, word := range words {
		var (
			counts = make(map[rune]int)
			count  int
		)

		for _, char := range word {
			counts[char]++
		}

		count = len(counts)
		if count > res {
			res = count
		}
	}

	return res
}
