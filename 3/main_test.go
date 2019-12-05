package main

import "testing"

func TestCalculateCabDistance(t *testing.T) {
	a := point{0, 0}
	b := point{1, 3}
	if a.taxiCabDistance(&b) != 4 {
		t.Fail()
	}
	if b.taxiCabDistance(&a) != 4 {
		t.Fail()
	}
}
