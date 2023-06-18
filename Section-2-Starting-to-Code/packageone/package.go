package packageone

import "fmt"

// Notice the capitalization of the variable names.
// This is how we export variables and functions from a package.
// var privateVar = "I am private"
// var PublicVar = "I am public (or exported)"

// func notExported() {

// }

// func Exported() {

// }

var PackageVar = "This is a package level variable in packageone."

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)
}