/*
2. The given program accepts 2 values from the user, length and breadth of a rectangle respectively.

Complete the program by writing methods "Area" and "Perimeter" on Rectangle type to calculate the area and perimeter of a rectangle.

Hint:
Method Signatures for Area and Perimeter
Area() int
Perimeter() int

Formulae:
Area = length * width
Perimeter = 2 * (length + width)

Example Test Case 1:
Input: 10 20
Output:
Area of Rectangle: 200
Perimeter of Rectangle: 60
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Rectangle struct {
	length, width int
}

func (rect Rectangle) Area() int {
	return rect.length * rect.width
}

func (rect Rectangle) Perimeter() int {
	return 2 * (rect.length + rect.width)
}

func ReadIntFromTerminal(message string) int {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic("Error while reading input")
	}

	// Convert the input to int
	input = strings.Trim(input, "\n")
	optionNumber, err := strconv.Atoi(input)
	if err != nil {
		panic("Please enter a valid integer")
	}

	return optionNumber
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	length := ReadIntFromTerminal("Enter rectangle length: ")
	width := ReadIntFromTerminal("Enter rectangle width: ")

	var rect = Rectangle{length, width}
	if rect.length < 0 || rect.width < 0 {
		fmt.Println("length or width cannot be negative.")
		return
	}

	fmt.Println("\nArea of Rectangle:", rect.Area())
	fmt.Println("Perimeter of Rectangle:", rect.Perimeter())
}
