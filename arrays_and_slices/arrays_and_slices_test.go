package arrays_and_slices

import (
	"reflect"
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

func TestSumAll(t *testing.T) {
	t.Run("return slice od sums for each passed slice", func(t *testing.T) {
		got := SumAll([]int{7, 2, 6}, []int{2, 3, 6, 8, 1})
		want := []int{15, 20}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\nWant: %d\nGot: %d", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, want, got []int) {
		t.Helper()
		for i, sum := range want {
			if sum != got[i] {
				t.Errorf("\nWant: %+v\nGot: %v", got, want)
			}
		}
	}
	t.Run("sums tails of all given slices", func(t *testing.T) {
		got := SumAllTails([]int{3, 2, 5, 3}, []int{2, 7, 9, 3})
		want := []int{10, 19}

		checkSums(t, want, got)
	})

	t.Run("summs empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 7, 9, 3})
		want := []int{0, 19}

		checkSums(t, want, got)
	})
}
