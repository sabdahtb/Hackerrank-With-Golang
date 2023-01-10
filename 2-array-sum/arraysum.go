package arraysum

func ArraySum(ar []int32) int32 {
	var result = int32(0)
	for _, s := range ar {
		result += s
	}
	return result
}
