package highestprofitwins

import "testing"

func TestMain(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		want := [2]int{1, 5}
		got := MinMax([]int{1, 2, 3, 4, 5})
		if want != got {
			t.Errorf("want %v got %v", want, got)
		}
	})

	t.Run("Test 2", func(t *testing.T) {
		want := [2]int{5, 2334454}
		got := MinMax([]int{2334454, 5})
		if want != got {
			t.Errorf("want %v got %v", want, got)
		}
	})

	t.Run("Test 3", func(t *testing.T) {
		want := [2]int{1, 1}
		got := MinMax([]int{1})
		if want != got {
			t.Errorf("want %v got %v", want, got)
		}
	})

}
