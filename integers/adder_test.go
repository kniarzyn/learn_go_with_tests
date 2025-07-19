package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)

	assert.Equal(t, 5, sum)
}

func ExampleAdd() {
	sum := Add(3, 5)
	fmt.Println(sum)
	// Output: 8
}
