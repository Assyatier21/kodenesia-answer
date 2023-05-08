package service

import (
	"strconv"
	"strings"
)

func CancelledSum(s string) int {
	var (
		res       int = 0
		lastCount int = 0
	)

	nums := strings.Split(s, " ")

	for _, num := range nums {
		if num == "B" {
			res = lastCount
		} else {
			val, _ := strconv.Atoi(num)
			lastCount = res
			res += val
		}
	}

	return res
}
