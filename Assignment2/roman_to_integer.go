/*
	Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000

	For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
	Roman numerals are usually written from largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.
	Given a roman numeral, convert it to an integer.

	Example 1:
	Input: s = "III"
	Output: 3
	Explanation: III = 3.

	Example 2:
	Input: s = "LVIII"
	Output: 58
	Explanation: L = 50, V= 5, III = 3.

	Example 3:
	Input: s = "MCMXCIV"
	Output: 1994
	Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

package main

import (
	"errors"
	"fmt"
)

func getValueFromSymbol(symbol byte) (int, error) {
	switch symbol {
	case 'I':
		return 1, nil
	case 'V':
		return 5, nil
	case 'X':
		return 10, nil
	case 'L':
		return 50, nil
	case 'C':
		return 100, nil
	case 'D':
		return 500, nil
	case 'M':
		return 1000, nil
	default:
		return -1, errors.New("Invalid Roman Symbol")
	}
}

func isSubtractSymbolCase(i int, s string) bool {
	var currentCh byte = s[i]
	var nextCh byte = '#'
	if i+1 < len(s) {
		nextCh = s[i+1]
	}

	if currentCh == 'I' && (nextCh == 'V' || nextCh == 'X') {
		return true
	} else if currentCh == 'X' && (nextCh == 'L' || nextCh == 'C') {
		return true
	} else if currentCh == 'C' && (nextCh == 'D' || nextCh == 'M') {
		return true
	} else {
		return false
	}
}

func romanToInt(s string) (int, error) {
	num := 0

	for i := range s {
		val, err := getValueFromSymbol(s[i])
		if err != nil {
			return -1, err
		}

		// Checking if we have to subtract or not based on the I, X, C condition
		if isSubtractSymbolCase(i, s) {
			num -= val
		} else {
			num += val
		}
	}

	return num, nil
}

func main() {
	var roman string
	fmt.Print("Enter roman numerals: ")
	_, err := fmt.Scanln(&roman)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	integer, err := romanToInt(roman)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Roman: ", roman)
	fmt.Println("Integer:", integer)
}
