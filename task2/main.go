package main

import (
	"github.com/umi3uwu/Assignment1_go/shapes"
)

func main() {
	shapesList := []shapes.Shape{
		shapes.Rectangle{Length: 10, Width: 5},
		shapes.Circle{Radius: 7},
		shapes.Square{Length: 4},
		shapes.Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for _, shape := range shapesList {
		shapes.PrintShapeDetails(shape)
	}
}
