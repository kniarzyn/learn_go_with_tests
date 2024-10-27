package arrays_and_slices

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{2, 5, 7, 9, 10}
		got := Sum(numbers)
		want := 33

		if want != got {
			t.Errorf("got: %d, want: %d when given %v", got, want, numbers)
		}
	})
}
