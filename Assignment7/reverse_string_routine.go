/*
	2. Given a string, reverse it using one go routine.And inside go routine print reversed string and number of goroutines launched
	e.g: Input: test123 output: 321tset 2
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

func reverseString(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	reversedString := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversedString += string(s[i])
	}

	fmt.Printf("%s %d", reversedString, runtime.NumGoroutine())
}

func main() {
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	input = strings.Trim(input, "\n")

	wg := sync.WaitGroup{}

	wg.Add(1)
	go reverseString(input, &wg)

	wg.Wait()
}
