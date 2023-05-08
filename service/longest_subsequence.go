package service

func LongestSequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	BubbleSort(arr)
	seqCounter, maxSeqLen := 1, 1

	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1]+1 {
			seqCounter++
		} else {
			seqCounter = 1
		}

		if seqCounter > maxSeqLen {
			maxSeqLen = seqCounter
		}
	}

	return maxSeqLen
}
