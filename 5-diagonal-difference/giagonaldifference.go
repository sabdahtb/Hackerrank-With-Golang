package diagonaldifference

import "math"

func DiagonalDifference(arr [][]int32) int32 {
	var a = int32(0)
	var b = int32(0)
	var length = len(arr) - 1
	for i := 0; i < len(arr); i++ {
		a += arr[i][i]
		b += arr[i][length]
		length -= 1
	}
	return int32(math.Abs(float64(a) - float64(b)))
}
