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
	"fmt"
	"log"
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

func main() {
	var index int
	fmt.Print("Enter an integer: ")
	_, err := fmt.Scanln(&index)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println(getDayFromInteger(index))
}
