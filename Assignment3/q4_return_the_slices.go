/*
4. Return the slices
Complete the program to return 3 slices of a given array, based on the following conditions.
Given Array : ["qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"]
Hint: Hard-code the given array into your program.

Input: Two space-separated integers representing index1 and index2.
Output: Output will contain 3 slices
1. slice containing all the elements from the start of the array to Index1
2. slice containing all the elements from index1 to index2
3. slice containing all the elements from index2 to the end of the array
Conditions to Handle:
If Either of the input indexes is out of range of the given array, print "Incorrect Indexes" and exit the program

Example Test Case 1:
Input: 2 4
Output:
[qwe wer ert]
[ert rty tyu]
[tyu yui uio iop]
Example Test Case 2:
Input: 2 8
Output: Incorrect Indexes
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sliceArrayOnIndex(index1 int, index2 int) ([]string, []string, []string) {
	// Given Hardcoded Array
	arr := [...]string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	// Exit as per constraint
	if index1 > index2 || index1 < 0 || index2 >= len(arr) {
		log.Fatalln("Incorrect Indexes")
	}

	slice1 := arr[:index1+1]
	slice2 := arr[index1 : index2+1]
	slice3 := arr[index2:]

	return slice1, slice2, slice3
}

func ReadMultiIntFromTerminal(message string, outputCount int) ([]int, error) {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return []int{}, errors.New("Error while reading input")
	}

	// Removing the last \n from the buffer
	input = strings.Trim(input, "\n")

	// convert the input to tokens and then read one by one
	tokens := strings.Fields(input)

	numbers := []int{}
	for _, token := range tokens {
		number, err := strconv.Atoi(token)
		if err != nil {
			return []int{}, errors.New("Please enter a valid integer")
		}

		numbers = append(numbers, number)
	}

	// Check if the numbers and count is equal
	if len(numbers) != outputCount {
		return []int{}, errors.New(fmt.Sprintf("Please only enter %d numbers as input", outputCount))
	}

	return numbers, nil
}

func main() {
	indexes, err := ReadMultiIntFromTerminal("Enter an integer: ", 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	index1, index2 := indexes[0], indexes[1]
	slice1, slice2, slice3 := sliceArrayOnIndex(index1, index2)

	// Printing slices on newlines
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
