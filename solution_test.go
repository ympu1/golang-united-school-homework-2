package square

import "testing"

const (
	sideLen = 10.0
)

func TestCalcSquare3Sides(t *testing.T) {
	got := CalcSquare(sideLen, 3)
	want := 43.30127018922193
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalcSquare4Sides(t *testing.T) {
	got := CalcSquare(sideLen, 4)
	want := 100.0
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalcSquare0Sides(t *testing.T) {
	got := CalcSquare(sideLen, 0)
	want := 314.1592653589793
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalcSquareWrongSides(t *testing.T) {
	var i customInt
	want := 0.0
	for i = -1000; i < 1000; i++ {
		if i == 0 || i == 3 || i == 4 {
			continue
		}
		got := CalcSquare(sideLen, i)
		if got != want {
			t.Errorf("got %v, wanted %v; SidesNum=%d", got, want, i)
		}
	}
}
