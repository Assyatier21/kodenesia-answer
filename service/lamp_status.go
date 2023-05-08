package service

import (
	"fmt"
	"strings"
)

func LampStatus(lamps int, switchs int) int {
	var (
		lampSwitch []int = make([]int, lamps)
		res        int   = 0
	)

	for i := range lampSwitch {
		lampSwitch[i] = 0
	}

	for i := 1; i <= switchs; i++ {
		for j := range lampSwitch {
			if (j+1)%i == 0 {
				switch lampSwitch[j] {
				case 0:
					lampSwitch[j] = 1
				case 1:
					lampSwitch[j] = 0
				}
			}
		}
	}

	res = strings.Count(fmt.Sprintf("%v", lampSwitch), "1")
	return res
}
