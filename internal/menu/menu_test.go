package menu

import "testing"

func TestPrice(t *testing.T) {
	if c, ok := Price("espresso"); !ok || c != 220 {
		t.Errorf("espresso = %d, %v", c, ok)
	}
	if _, ok := Price("hot dog"); ok {
		t.Error("we do not sell hot dogs")
	}
}
