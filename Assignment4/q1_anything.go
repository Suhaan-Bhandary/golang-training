/*
	1. The given program accepts an integer value between 1 to 4 from the user.
	There is an option associated with each value, which is basically objects of different data types with some associated value.

	Write a method named "AcceptAnything" which takes any type of data as input.

	Based on the option chosen by the user, we will send different types of objects to this function and it should print the following based on its respective data type of value sent to the function "AcceptAnything".

	integer :
	"This is a value of type Integer, <value>"

	string :
	"This is a value of type String, <value>"

	boolean :
	"This is a value of type Boolean, <value>"

	Custom data type Hello :
	"This is a value of type Hello, <value>"

	Input: 1,2,3,4

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

type Hello struct {
	name string
}

func AcceptAnything(value any) {
	switch assertedValue := value.(type) {
	case int:
		fmt.Println("This is a value of type Integer,", assertedValue)
	case string:
		fmt.Println("This is a value of type String,", assertedValue)
	case bool:
		fmt.Println("This is a value of type Boolean,", assertedValue)
	case Hello:
		fmt.Println("This is a value of type Hello,", assertedValue)
	default:
		fmt.Println("Unknown input types,", assertedValue)
	}
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
	optionNumber, err := ReadIntFromTerminal("Enter an index: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// ! We have informed to use hard coded value here
	switch optionNumber {
	case 1:
		AcceptAnything(5)
	case 2:
		AcceptAnything("Go")
	case 3:
		AcceptAnything(true)
	case 4:
		AcceptAnything(Hello{})
	default:
		fmt.Println("Option value should be between 1 to 4(both inclusive)")
	}
}
