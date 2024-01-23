/*
3. The given program takes an integer value as input from the user i.e 1 or 2.
Option 1 represents Rectangle and Option 2 represents Square.

Given the metrics of the shapes are hard-coded, complete the given program to achieve the following:

1. Create an interface Quadrilateral which has the following method signatures
Area() int
Perimeter() int

2. Implement the Quadrilateral interface for the given shapes i.e Rectangle and Square

3. Define a "Print" function which accepts any shape that implements Quadrilateral interface and Prints Area and Perimeter of shape in the following manner:

"Area :  <value>"
"Perimeter :  <value>"


HINT: Step 2 means, to define methods in Quadrilateral interface for given shapes


Formulae:

Area of Rectangle: length * width
Perimeter of Rectangle: 2*(length + breadth)


Area of Square: side * side
Perimeter of Square: 4 * side
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rectangle struct {
	length, width int
}

type Square struct {
	side int
}

type Quadrilateral interface {
	Area() int
	Perimeter() int
}

func (rect Rectangle) Area() int {
	return rect.length * rect.width
}

func (rect Rectangle) Perimeter() int {
	return 2 * (rect.length + rect.width)
}

func (square Square) Area() int {
	return square.side * square.side
}

func (square Square) Perimeter() int {
	return 4 * square.side
}

func Print(quad Quadrilateral) {
	fmt.Println("Area :", quad.Area())
	fmt.Println("Perimeter :", quad.Perimeter())
}

func ReadIntFromTerminal(message string) (int, error) {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, errors.New("Error while reading input")
	}

	// Convert the input to int
	input = strings.Trim(input, "\n")
	optionNumber, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("Please enter a valid integer")
	}

	return optionNumber, nil
}

func main() {
	optionNumber, err := ReadIntFromTerminal("Enter an Option 1/2: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	switch optionNumber {
	case 1:
		length, err := ReadIntFromTerminal("Enter rectangle length: ")
		if err != nil {
			fmt.Println(err)
			return
		}

		width, err := ReadIntFromTerminal("Enter rectangle width: ")
		if err != nil {
			fmt.Println(err)
			return
		}

		rect := Rectangle{length, width}
		Print(rect)
	case 2:
		side, err := ReadIntFromTerminal("Enter square side: ")
		if err != nil {
			fmt.Println(err)
			return
		}

		square := Square{side}
		Print(square)
	default:
		fmt.Println("Invalid Option, Option should be 1 or 2.")
	}

}
