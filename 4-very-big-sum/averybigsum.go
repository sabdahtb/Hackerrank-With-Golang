package verybigsum

func AVeryBigSum(ar []int64) int64 {
	var result = int64(0)
	for _, s := range ar {
		result += s
	}
	return result
}
