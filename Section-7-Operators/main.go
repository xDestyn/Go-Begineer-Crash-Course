package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// Operators and precedence
	answer := 7 + 3*4
	fmt.Println("Answer:", answer)

	answer = (7 + 3) * 4
	fmt.Println("Answer:", answer)

	// Primary operators
	// Multiplication
	var radius = 12.0
	area := radius * radius * math.Pi
	fmt.Println("Area is:", area)

	// Division
	half := 1 / 2
	fmt.Println("Half int is:", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("Half float is:", halfFloat)

	// Square
	badThreeSquared := 3 ^ 2
	fmt.Println("Bad three squared is:", badThreeSquared)

	goodThreeSquared := math.Pow(3, 2)
	fmt.Println("Good three squared is:", goodThreeSquared)

	// Modulus
	remainder := 7 % 3
	fmt.Println("Remainder is:", remainder)

	// Unary operators
	x := 3
	x++
	fmt.Println("X is:", x)
	x--
	fmt.Println("X is:", x)

	// Relational operators
	second := 31
	minute := 1

	if minute < 59 && second+1 > 59 {
		minute++
	}
	fmt.Println(minute)

	// Short circuiting
	a := 12
	b := 0

	// if b != 0 {
	// 	c := divideTwoNumbers(a, b)
	// 	if c == 2 {
	// 		fmt.Println("We found a two")
	// 	}
	// }

	// if b != 0 && divideTwoNumbers(a, b) == 2 {
	// 	fmt.Println("We found a two")
	// }

	c, err := divideTwoNumbers(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		if c == 2 {
			fmt.Println("We found a two")
		}
	}

	// Assignment operators
	o := 12
	o++
	fmt.Println("O is:", o)

	p := 10
	p /= 2
	fmt.Println("P is:", p)

}

func divideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return x / y, nil
	}
}
