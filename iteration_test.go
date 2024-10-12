package iteration

import "testing"

func TestRepeat(t *testing.T) {
	r := Repeat("n", 5)
	expected := "nnnnn"

	if r != expected {
		t.Errorf("expected %q - got %q", expected, r)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("s", 5)
	}
}
