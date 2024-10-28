package arrays_and_slices

func Sum(s []int) int {
	var sum int

	for _, n := range s {
		sum += n
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	var r []int
	for _, s := range slices {
		r = append(r, Sum(s))
	}
	return r
}

func SumAllTails(slices ...[]int) []int {
	var sums []int

	for _, s := range slices {
		el := 0
		if len(s) > 0 {
			el = Sum(s[1:])
		}
		sums = append(sums, el)
	}

	return sums
}
