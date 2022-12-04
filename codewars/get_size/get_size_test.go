package getsize

import "testing"

func TestGetSize(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		got := GetSize(4, 2, 6)
		want := [2]int{88, 48}
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
