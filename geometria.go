package main

import "math"

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func Distancia(p1, p2 Point) float64 {
	deltaX := p1.X - p2.X
	deltaY := p1.Y - p2.Y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
