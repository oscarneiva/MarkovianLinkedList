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

func (l *Line) calculateDistance() float64 {
	distance := math.Sqrt(math.Pow(l.p1.x-l.p2.x, 2) + math.Pow(l.p1.y-l.p2.y, 2))
	l.p1.x = 1000
	return distance
}

//func New(x1 float64, y1 float64, x2 float64, y2 float64) *Line{
//	return &Line{Point{x1,y1}, Point{x2,y2}}
//}

func main() {
	//l := New(1,1,2,2)
	l := Line{Point{1, 1}, Point{2, 2}}
	fmt.Println(l.calculateDistance()) // Use values in the constructor to calculate distance
	fmt.Println(l.calculateDistance()) // Print value after changing l.p1.x to 1000
}
