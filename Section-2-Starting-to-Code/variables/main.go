package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	// Variables
	// One way - Declare then assign (two steps)
	// var firstNumber int
	// firstNumber = 2
	// fmt.Println(firstNumber)

	// Another way - Declare type and name and assign value
	// var secondNumber = 5
	// fmt.Println(secondNumber)

	// One step variable - Declare name, assign value and let Go figure out the type
	// subtraction := 7
	// fmt.Println(subtraction)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber * secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
	
}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)
	
	// Display a welcome/instructions message
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// Take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// Give them the answer
	fmt.Println("The answer is", answer)
}