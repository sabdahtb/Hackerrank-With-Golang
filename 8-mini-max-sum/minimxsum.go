package minimaxsum

func MiniMaxSum(arr []int32) []int32 {
	var min = int32(arr[0])
	var max = int32(arr[0])
	var total = int32(0)

	for _, value := range arr {
		total += value

		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return []int32{total - min, total - max}
}
