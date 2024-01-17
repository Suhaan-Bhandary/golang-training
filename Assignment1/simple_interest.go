package main

import (
	"fmt"
)

func main() {
	var principle, rate, time float64 = 1000, 8.32789, 3

	interest := (principle * rate * time) / 100
	fmt.Println("Simple Interest:", fmt.Sprintf("%.2f", interest))
}
