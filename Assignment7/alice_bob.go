/*
	1. Given a string containing conversation between alice and bob. In the string, if it reaches $, it means end of alice message and if it reaches #, it means end of bob's message. and if it reaches ^,it means end of conversation ignore string after that.

	Note: given string doesn't contain any spaces. If message contains two continuous conversations from one person it should be printed one after another as given in the example. write a program to separate out messages from alice and bob. Write messages from alice and bob on seperate channels, whenever a message from alice/bob, print it in front of their name as shown in the example below:

	Note: there is a single space before and after colon(:) and no space before and after comma.

	e.g: "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"
	output: alice : helloBob,bob : helloalice,bob : howareyou?,alice : Iamgood.howareyou?
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func writeToRespectiveChannels(input string, aliceChan chan<- string, bobChan chan<- string, signalChan chan<- struct{}) {
	defer func() {
		signalChan <- struct{}{}
	}()

	current := ""
	for _, ch := range input {
		switch ch {
		case '$':
			aliceChan <- current
			current = ""
		case '#':
			bobChan <- current
			current = ""
		case '^':
			return
		default:
			current += string(ch)
		}
	}
}

func processInputConversation(input string) error {
	if len(input) == 0 || input[len(input)-1] != '^' {
		return errors.New("Error: Input conversation should end with ^")
	}

	if len(input) > 1 && input[len(input)-2] != '$' && input[len(input)-2] != '#' {
		return errors.New("Error: Input conversation should have # or $ before ^")
	}

	// Storing alice and bob messages and signalChan is used to tell that function has completed
	aliceChan := make(chan string)
	bobChan := make(chan string)
	signalChan := make(chan struct{})

	defer func() {
		close(aliceChan)
		close(bobChan)
		close(signalChan)
	}()

	go writeToRespectiveChannels(input, aliceChan, bobChan, signalChan)

	for {
		select {
		case val := <-aliceChan:
			fmt.Printf("alice : %s\n", val)
		case val := <-bobChan:
			fmt.Printf("bob : %s\n", val)
		case <-signalChan:
			return nil
		}
	}
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

	err = processInputConversation(input)
	if err != nil {
		fmt.Println(err)
	}
}
