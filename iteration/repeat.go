package iteration

import (
	"fmt"
	"strings"
)

// Function repeats given string count times.
func Repeat(char string, count int) string {
	var res strings.Builder

	for range count {
		res.WriteString(char)
	}
	return res.String()
}

func main() {
	fmt.Println(Repeat("s", 5))
}
