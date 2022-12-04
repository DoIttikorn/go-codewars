package convertstringarray

import (
	"strings"
)

func StringToArray(str string) []string {
	return strings.Fields(str)
}
