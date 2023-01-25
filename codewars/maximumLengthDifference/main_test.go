package maximumlengthdifference

import (
	"testing"
)

func TestMxDifLg(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		want := 13
		got := MxDifLg([]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfq knn", "qqquuhii", "dvvvwz"}, []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"})
		if want != got {
			t.Errorf("want %v got %v", want, got)
		}
	})

	// t.Run("Test 2", func(t *testing.T) {
	// 	want := -1
	// 	got := MxDifLg([]string{"it", "wkppv", "ixoyx", "3452", "zzzzzzzz zz"}, []string{"x", "3452", "zzzzzzzz zz"})
	// 	if want != got {
	// 		t.Errorf("want %v got %v", want, got)
	// 	}
	// })

	t.Run("Test 3", func(t *testing.T) {
		want := 10
		got := MxDifLg([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, []string{"bbbaaayddqbbrrrv"})
		if want != got {
			t.Errorf("want %v got %v", want, got)
		}
	})
}
