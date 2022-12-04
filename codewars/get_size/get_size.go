package getsize

func GetSize(w, h, d int) [2]int {
	return [2]int{2 * (w*h + h*d + w*d), w * h * d}
}
