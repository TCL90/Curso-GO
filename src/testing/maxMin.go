package maxMin

func Max(sliceInts []int) int {
	if len(sliceInts) == 0 {
		return -1
	}
	var maxInt int = sliceInts[0]
	for _, v := range sliceInts {
		if v > maxInt {
			maxInt = v
		}
	}
	return maxInt
}

func Min(sliceInts []int) int {
	if len(sliceInts) == 0 {
		return -1
	}
	var minInt int = sliceInts[0]
	for _, v := range sliceInts {
		if v < minInt {
			minInt = v
		}
	}
	return minInt
}
