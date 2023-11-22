package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello World"
	fmt.Println("String:", name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRune\tString")
	for index, rune := range name {
		fmt.Println(index, "\t", rune, "\t", string(rune))
	}
	fmt.Println()

	name = "Γειά σου Κόσμε"
	fmt.Println("Index\tRune\tString")
	for index, rune := range name {
		fmt.Println(index, "\t", rune, "\t", string(rune))
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("Three ways to concatenate strings")

	h := "Hello, "
	w := "World!"

	// using + not very efficient
	myString := h + w
	fmt.Println(myString)
	fmt.Println()

	// using fmt - more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)
	fmt.Println()

	// using stringbuilder - very effecient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()

	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Getting a substring")
	fmt.Println(name[0:13])

	// Indexes
	courseName := "Learn Go for Beginners Crash Course"

	fmt.Println(courseName[0])
	fmt.Println(courseName[6])

	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()

	for i := 13; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()

	fmt.Println("Length of courseName:", len(courseName))

	var mySlice []string
	mySlice = append(mySlice, "One")
	mySlice = append(mySlice, "Two")
	mySlice = append(mySlice, "Three")

	fmt.Println("mySlice:", len(mySlice), "elements")
	fmt.Println("The last element is:", mySlice[len(mySlice)-1])

	// Strings package
	courses := []string{
		"Learn Go for Beginners Crash Course",
		"Learn Docker and Kubernetes Crash Course",
		"Learn Python for Beginners Crash Course",
		"Learn Javascript and React for Beginners Crash Course",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in:", x, "at index:", strings.Index(x, "Go"))
		}
	}

	newString := "Go is an awesome language. Go for it!"
	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasSuffix(newString, "Python"))
	fmt.Println(strings.HasSuffix(newString, "it!"))
	fmt.Println(strings.Count(newString, "Go"))
	fmt.Println(strings.Count(newString, "Python"))
	fmt.Println(strings.Index(newString, "Python"))
	fmt.Println(strings.LastIndex(newString, "Go"))

	// String manipulation
	newString2 := "Go is an awesome language. Go for it!"

	if strings.Contains(newString2, "Go") {
		// newString2 = strings.Replace(newString2, "Go", "Golang", 1)
		newString2 = strings.ReplaceAll(newString2, "Go", "Golang")
	}

	fmt.Println(newString2)

	// String comparison
	a := "A"
	if a == "A" {
		fmt.Println("They are equal")
	}

	if "A" > "B" {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println("B is greater than A")
	}

	badEmail := "me@example.com  "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Printf("=%s=", badEmail)
	fmt.Println()

	str := "alpha alpha alpha alpha alpha alpha"
	str = replaceNth(str, "alpha", "beta", 3)
	fmt.Println(str)

	myStr := "This is a clear example of why we search in one case only."
	searchString := strings.ToLower(myStr)
	if strings.Contains(searchString, "this") {
		fmt.Println("This is found")
	} else {
		fmt.Println("This is not found")
	}

	fmt.Println(strings.ToLower(searchString))
	fmt.Println(strings.ToUpper(searchString))
	fmt.Println(strings.Title(strings.ToLower(searchString)))
}

func replaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i += x
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}
