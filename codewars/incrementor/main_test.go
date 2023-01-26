package incrementor

import "testing"

func TestIncrementor(t *testing.T) {
	t.Run("Basic tests", func(t *testing.T) {
		var tests = []struct {
			input []int
			want  []int
		}{
			{[]int{1, 2, 3}, []int{2, 4, 6}},
			{[]int{4, 6, 7, 1, 3}, []int{5, 8, 0, 5, 8}},
			{[]int{3, 6, 9, 8, 9}, []int{4, 8, 2, 2, 4}},
			{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 9, 8}, []int{2, 4, 6, 8, 0, 2, 4, 6, 8, 9, 0, 1, 2, 2}},
			{[]int{1}, []int{2}},
			{[]int{}, []int{}},
		}
		for _, tt := range tests {
			got := Incrementor(tt.input)
			if !equal(got, tt.want) {
				t.Errorf("Incrementor(%v) = %v, want %v", tt.input, got, tt.want)
			}
		}
	})
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
