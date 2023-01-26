package makeafunctionthatdoesarithmetic

func Arithmetic(a, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	default:
		panic("Unknow operator")
	}
}
