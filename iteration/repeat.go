package iteration

import "strings"

func Repeat(char string, count int) string {
	var res strings.Builder

	for range count {
		res.WriteString(char)
	}
	return res.String()
}
