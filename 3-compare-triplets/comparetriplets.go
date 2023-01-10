package comparetriplets

func CompareTriplets(a []int32, b []int32) []int32 {
	var pointA = int32(0)
	var pointB = int32(0)

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			pointA += 1
		}

		if a[i] < b[i] {
			pointB += 1
		}
	}

	return []int32{pointA, pointB}
}
