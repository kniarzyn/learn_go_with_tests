package arrays_and_slices

func Sum(s []int) int {
	var sum int

	for _, n := range s {
		sum += n
	}
	return sum
}
