package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 8.0, height: 4.0}
	got := Perimeter(rectangle)
	want := 24.0

	if float64(got) != want {
		t.Errorf("For 4 and 9 want: %.2f; got: %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	testCases := []struct {
		shape Shape
		want  float64
	}{
		{Circle{10}, 314.1592653589793},
		{Rectangle{8, 4}, 32.00},
	}

	for _, tc := range testCases {
		got := tc.shape.Area()
		if float64(got) != tc.want {
			t.Errorf("for shape: %#v -> want: %.2f; got: %.2f", tc.shape, tc.want, got)
		}
	}
}
