package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Line struct {
	p1 Point
	p2 Point
}

func (l Line) calculateDistance() float64{
	return math.Sqrt(math.Pow(l.p1.x - l.p2.x,2) + math.Pow(l.p1.y - l.p2.y,2))
}

func main() {
	l := Line{Point{1,1}, Point{2,2}}
	fmt.Println(l.calculateDistance())
}




