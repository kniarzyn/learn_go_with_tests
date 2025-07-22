package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 5)

	assert.Equal(t, "aaaaa", result)
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 7)
	fmt.Println(res)
	// Output: aaaaaaa
}
