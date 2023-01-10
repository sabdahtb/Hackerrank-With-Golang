package birthdaycakecandles

func BirthdayCakeCandles(candles []int32) int32 {
	var totalMax = int32(0)

	var a = candles[0]

	for _, s := range candles {
		if s > a {
			a = s
		}
	}

	for _, t := range candles {
		if t == a {
			totalMax += 1
		}
	}

	return totalMax
}
