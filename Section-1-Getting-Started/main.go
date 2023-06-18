package main

import (
	"Section-1-Getting-Started/doctor"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1) Way to declare and initialize a variable
	// var whatToSay string
	// whatToSay = "Hello, world!"

	// 2) Shorthand notation
	// whatToSay := "Hello, world!"
	// sayHelloWorld(whatToSay)

	// var whatToSay string

	// Read input from the user
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	// Only one loop in Go - For
	for {
		fmt.Print("-> ")
		// Capture input until user hits enter
		userInput, _ := reader.ReadString('\n')
		// fmt.Println(userInput)

		// Strip off '\n' character for Windows users
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		// Strip off '\n' character for Mac and Linux users
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			// response := doctor.Response(userInput)
			// fmt.Println(response)
			// Fewer lines of code, less to maintain!
			fmt.Println(doctor.Response(userInput))
		}
	}
}

// func sayHelloWorld(whatToSay string) {
// 	fmt.Println(whatToSay)
// }
