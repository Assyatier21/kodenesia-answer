package service

import (
	"strconv"
	"strings"
)

func FlyingRugs(s string) int {
	var (
		limit      int
		lastHeight int
		mtHeight   = []int{}
	)

	sArr := strings.Split(s, " ")
	limit, _ = strconv.Atoi(sArr[0])

	for i := 1; i < len(sArr); i++ {
		temp, _ := strconv.Atoi(sArr[i])
		mtHeight = append(mtHeight, temp)
	}

	for _, v := range mtHeight {
		if (lastHeight + limit) >= v {
			if v > lastHeight {
				lastHeight = v
			}
		}
	}

	return lastHeight
}
