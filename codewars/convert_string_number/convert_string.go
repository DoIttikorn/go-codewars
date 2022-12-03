package convertstringnumber

import "strconv"

func StringToNumber(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}
