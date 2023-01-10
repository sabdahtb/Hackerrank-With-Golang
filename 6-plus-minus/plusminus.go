package plusminus

import "fmt"

func PlusMinus(arr []int32) []string {
	var positive = int(0)
	var negative = int(0)
	var zero = int(0)
	for _, s := range arr {
		if s > 0 {
			positive += 1
		} else if s < 0 {
			negative += 1
		} else {
			zero += 1
		}
	}
	return []string{
		fmt.Sprintf("%.6f", float64(positive)/float64(len(arr))),
		fmt.Sprintf("%.6f", float64(negative)/float64(len(arr))),
		fmt.Sprintf("%.6f", float64(zero)/float64(len(arr)))}
}
