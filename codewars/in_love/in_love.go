package convertstringnumber

func LoveFunc(flower1, flower2 int) bool {
	if flower1%2 == 0 && flower2%2 == 1 {
		return true
	} else if flower1%2 == 1 && flower2%2 == 0 {
		return true
	} else {
		return false
	}
	// return (a+b) % 2 == 1 ||  (flower1 + flower2) %2 != 0
}

/*

Timmy & Sarah think they are in love, but around where they live,
they will only know once they pick a flower each. If one of the flowers
has an even number of petals and the other has an odd number of petals it means they are in love.


*/
