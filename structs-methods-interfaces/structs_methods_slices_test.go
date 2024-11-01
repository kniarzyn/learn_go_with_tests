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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		if float64(got) != want {
			t.Errorf("for shape: %#v -> want: %.2f; got: %.2f", shape, want, got)
		}

	}
	t.Run("for rectangle", func(t *testing.T) {
		rectangle := Rectangle{width: 8.0, height: 4.0}
		checkArea(t, rectangle, 32.00)
	})

	t.Run("for circle", func(t *testing.T) {
		circle := Circle{radius: 10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
