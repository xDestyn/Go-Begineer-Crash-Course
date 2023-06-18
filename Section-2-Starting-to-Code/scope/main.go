package main

import "Section-2-Starting-to-Code/packageone"

// var one = "One"
var myVar = "This is a package level variable."

func main() {
	// var one = "This is a block level variable"

	// fmt.Println(one)

	// myFunc()

	// newString := packageone.PublicVar
	// fmt.Println("From packageone: ", newString)

	// packageone.Exported()
	
	// <---------------------------------------------->

	// Variables
	// Declare a package level variable for the main
	// package named myVar

	// Declare a block variable for the main function
	// called blockVar
	var blockVar = "This is a block level variable."

	// Declare a package level variable in the packageone
	// package named PackageVar

	// Create an exported function in packageone called PrintMe

	// In the main function, print out the values of myVar,
	// blockVar, and PackageVar one one line
	packageone.PrintMe(myVar, blockVar)
}

// func myFunc() {
// 	// var one = "the number one"

// 	fmt.Println(one)
// }