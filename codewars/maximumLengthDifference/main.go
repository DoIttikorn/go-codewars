package maximumlengthdifference

// https://www.codewars.com/kata/5663f5305102699bad000056/train/go

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	var max int
	for _, v1 := range a1 {
		for _, v2 := range a2 {
			if len(v1) > len(v2) {
				if len(v1)-len(v2) > max {
					max = len(v1) - len(v2)
				}
			} else {
				if len(v2)-len(v1) > max {
					max = len(v2) - len(v1)
				}
			}
		}
	}
	return max
}
