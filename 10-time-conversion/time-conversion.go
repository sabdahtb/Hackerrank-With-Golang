package timeconversion

import (
	"fmt"
	"strconv"
	"strings"
)

func TimeConversion(s string) string {
	s = strings.ToLower(s)
	pm := strings.HasSuffix(s, "pm")
	am := strings.HasSuffix(s, "am")

	t := s[:len(s)-2]

	timeArr := strings.Split(t, ":")
	hours, minutes, seconds := timeArr[0], timeArr[1], timeArr[2]
	hourInt, _ := strconv.Atoi(hours)

	if am && hourInt == 12 {
		hourInt = 0
	}

	if pm && hourInt != 12 {
		hourInt += 12
	}

	hours = strconv.Itoa(hourInt)

	return fmt.Sprintf("%02s:%02s:%02s", hours, minutes, seconds)
}
