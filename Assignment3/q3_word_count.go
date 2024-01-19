/*
3. Word Count
Write a Program to fulfil the following conditions.

Input: A string containing words separated by space
Output: A slice containing words with the highest frequency in the input string.
Note: The order of the words in the resultant slice should be the same as the order they appear in the input string.
Condition: Words are case-sensitive. Joe is not the same as joe.
Given Code Template: A code template is provided that takes a string as an input and turns it into a slice of strings.

Example Test Case 1:
Input: My name is Joe and My Father's name is also Joe
Output: [My name is Joe]
Here, the words "My", "name", "is", "Joe" appeared 2 times each, which is also the highest frequency of any word.
Hence the output contains all 4 words.

Example Test Case 2:
Input: Europe is far but the US is farther
Output: [is]
Here, the word "is" appeared twice which is also the highest frequency of any word.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getHighestFrequencyWords(sentence string) []string {
	words := strings.Split(sentence, " ")

	// Finding frequency of each word
	maxFreq := 0
	freq := map[string]int{}
	for _, word := range words {
		freq[word]++
		maxFreq = max(freq[word], maxFreq)
	}

	var highestFrequencyWords []string
	for _, word := range words {
		if freq[word] == maxFreq {
			highestFrequencyWords = append(highestFrequencyWords, word)
			freq[word] = -1
		}
	}

	return highestFrequencyWords
}

func main() {
	// Reading whole line from the user and ignoring \n at last
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)
	sentence, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// Removing the extra \n and " "
	sentence = strings.Trim(sentence, " \n")

	highestFrequencyWords := getHighestFrequencyWords(sentence)
	fmt.Println(highestFrequencyWords)
}
