package spatial

import "github.com/golang/geo/r2"

type Circle struct {
	Location r2.Point
	Radius   float32
}

func NewCircle() *Circle {
	var circle *Circle = new(Circle)

	return circle
}
