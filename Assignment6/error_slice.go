/*
	2. In the given code, the accessSlice function accepts a slice and index.
	If the value is present in the slice at that index, the program should print the following.
	"item <index>, value <value present at the index>"

	But if the index does not hold any value, it will lead to an index out of range panic in our program.

	So in order to safeguard our program from panicking, add a condition to handle the condition if
	index > lengthOfSlice - 1 and return an error from the accessSlice function if the above condition is met.
	The error message should be the following : "length of the slice should be more than index"

	Complete the given program to return an error from the accessSlice function and handle the returned error inside the main function to print the message.

	Example Test Case 1 :
	Input: 3
	Output:
	item 3, value 6
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

func accessSlice(slice []int, index int) error {
	if index > len(slice)-1 {
		return errors.New("length of the slice should be more than index")
	}

	fmt.Printf("item %d, value %d", index, slice[index])
	return nil
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

	err = accessSlice(slice, index)
	if err != nil {
		fmt.Println(err)
	}
}
