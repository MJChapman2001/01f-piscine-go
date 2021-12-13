package main

import "fmt"

type point struct {
	x int
	y int
}

func setPoint(prt *point) {
	prt.x = 42
	prt.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
