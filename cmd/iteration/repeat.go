package iteration

import (
	"fmt"
	"iter"
	"strings"
	"unicode"
)

// Function repeats given string count times.
func Repeat(char string, count int) string {
	var res strings.Builder

	for range count {
		res.WriteString(char)
	}
	return res.String()
}

func FunWithFieldsFunc(s string) iter.Seq[string] {
	return strings.FieldsFuncSeq(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}

func main() {
	fmt.Println(Repeat("s", 5))
}
