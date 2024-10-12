package iteration

// Repeats given string n times
func Repeat(c string, n int) string {
	var result string

	for i := 0; i < n; i++ {
		result += c
	}
	return result
}
