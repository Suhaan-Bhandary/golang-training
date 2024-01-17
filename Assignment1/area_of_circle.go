package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64 = 123

	area := math.Pi * math.Pow(radius, 2)
	fmt.Println("Area of circle:", fmt.Sprintf("%.2f", area))
}
