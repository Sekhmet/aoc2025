package toolkit

import "testing"

func TestPmod(t *testing.T) {
	got := Pmod(-1, 5)

	if got != 4 {
		t.Errorf("Pmod(-1, 5) = %d; want 4", got)
	}

	got = Pmod(6, 5)
	if got != 1 {
		t.Errorf("Pmod(6, 5) = %d; want 1", got)
	}
}
