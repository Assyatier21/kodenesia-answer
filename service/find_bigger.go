package service

import (
	"strconv"
)

var (
	counter  = 0
	givenNum int
)

func FindBigger(num int) int {
	var (
		nums []rune
	)
	s := strconv.Itoa(num)
	nums = []rune(s)

	givenNum = num

	permutations(nums, 0, len(nums)-1)
	return counter
}

func permutations(runes []rune, left, right int) {
	if left == right {
		k, _ := strconv.Atoi(string(runes))
		if k > givenNum {
			counter++
		}
	} else {
		for i := left; i <= right; i++ {
			runes[left], runes[i] = runes[i], runes[left]
			permutations(runes, left+1, right)
			runes[left], runes[i] = runes[i], runes[left]
		}
	}
}
