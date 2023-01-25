package maximumlengthdifference

import "math"

func findShortestAndLongest(a []string) (shortest, longest int) {
	lA := 0
	sA := math.MaxInt64
	for _, s := range a {
		if len(s) > lA {
			lA = len(s)
		}
		if len(s) < sA {
			sA = len(s)
		}
	}

	return sA, lA
}

func MxDifLg2(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	sA1, lA1 := findShortestAndLongest(a1)
	sA2, lA2 := findShortestAndLongest(a2)

	res1 := math.Abs(float64(sA1 - lA2))
	res2 := math.Abs(float64(sA2 - lA1))

	if res1 > res2 {
		return int(res1)
	}

	return int(res2)
}
