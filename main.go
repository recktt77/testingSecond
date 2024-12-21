package main

import (
	"fmt"
	"github.com/NameSurname/Assignment1/Shapes"
)

func PrintShapeDetails(s Shapes.Shape) {
	fmt.Printf("Shape Details:\n")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println()
}

func main() {
	shapes := []Shapes.Shape{
		Shapes.Rectangle{Length: 5, Width: 3},
		Shapes.Circle{Radius: 4},
		Shapes.Square{Length: 2.5},
		Shapes.Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for _, shape := range shapes {
		PrintShapeDetails(shape)
	}
}
