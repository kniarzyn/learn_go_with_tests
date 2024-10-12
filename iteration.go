package iteration

func Repeat(c string, r int) string {
	var result string

	for i := 0; i < r; i++ {
		result = result + c
	}
	return result
}
