/*
	1. In the given code, the accessSlice function accepts a slice and index.
	If the value is present in the slice at that index, the program should print the following.

	"item <index>, value <value at that index in slice>"

	But if the index does not hold any value,
	it will lead to an index out of range panic in our program.

	Complete the given code to recover from panic inside the accessSlice function, when index out of range panic occurs.
	Also, Print the following after handling panic

	"internal error: <recovered panic value>"


	Example Test Case 1 :

	Input: 3
	Output:
	item 3, value 6
	Testing Panic and Recover
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

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("internal error:", r)
		}
	}()

	fmt.Printf("item %d, value %d", index, slice[index])
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
	slice := []int{1, 2, 3, 4, 5}
	index, err := ReadIntFromTerminal("Enter an index: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	accessSlice(slice, index)
}
