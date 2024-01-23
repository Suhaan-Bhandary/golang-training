/*
	2. What's The Day
	Write a program to store the day(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday) against the respective index of the day(1, 2, 3, 4, 5, 6, 7) in a map.
	You can hard-code the map in your program.
	Take an integer input from the user and print the day stored against that index and if nothing is to be found for that index,
	Print "Not a day"
	Hint: Following code can be used to take an integer input from the user in the GO Programming Language
	var index int
	fmt.Scanln(&index)
	Example Test Case:
	Input: 5
	Output: Friday
	Input: 11
	Output: Not a day
*/

/*
	Same Problem statement as Question 1
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

// Returns Weekday for a given integer if present in map else "Not a Day"
func getDayFromInteger(index int) string {
	var weekDays = map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	weekDay, isPresent := weekDays[index]
	if !isPresent {
		return "Not a Day"
	}

	return weekDay
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
	index, err := ReadIntFromTerminal("Enter an integer: ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(getDayFromInteger(index))
}
