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
