package main

import "testing"

func TestDistancia(t *testing.T) {
	p1 := Point{X: 0, Y: 0}
	p2 := Point{X: 3, Y: 4}
	d := Distancia(p1, p2)

	if d != 5 {
		t.Error("a distância deveria ser 5")
	}
}
