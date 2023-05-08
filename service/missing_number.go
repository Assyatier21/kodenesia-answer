package service

import (
	"math"
)

func getValue(str string, i, m int) int {
	var (
		value = 0
	)
	if i+m > len(str) {
		return -1
	}

	for j := 0; j < m; j++ {
		c := int(str[i+j] - '0')
		if c < 0 || c > 9 {
			return -1
		}
		value = value*10 + c
	}

	return value
}
func MissingNumber(str string) int {
	var (
		res       int  = -1
		maxDigits int  = 7
		fail      bool = false
	)

	for m := 1; m <= maxDigits; m++ {
		n := getValue(str, 0, m)
		if n == -1 {
			break
		}

		for i := m; i != len(str); i += 1 + int(math.Log10(float64(n))) {
			if res == -1 && getValue(str, i, 1+int(math.Log10(float64(n+2)))) == n+2 {
				res = n + 1
				n += 2
			} else if getValue(str, i, 1+int(math.Log10(float64(n+1)))) == n+1 {
				n++
			} else {
				fail = true
				break
			}
		}

		if !fail {
			return res
		}
	}

	return -1
}

