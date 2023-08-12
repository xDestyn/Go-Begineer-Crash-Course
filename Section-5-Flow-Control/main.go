package main

import (
	"fmt"
)

func main() {
	// For 3 part loop
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("i is", i)
	// }

	// for i := 10; i >= 0; i-- {
	// 	fmt.Println("i is", i)
	// }

	// // While loop
	// rand.Seed(time.Now().UnixNano())
	// i := 1000
	// for i > 100 {
	// 	i = rand.Intn(1000) + 1
	// 	fmt.Println("i is", i)
	// 	if i > 100 {
	// 		fmt.Println("i is", i, "so loop keeps going")
	// 	}
	// }

	// fmt.Println("Got", i, "and broke out of loop")

	// Infinite Loop
	// reader := bufio.NewReader(os.Stdin)
	// ch := make(chan string)

	// go mylogger.ListenForLog(ch)

	// fmt.Println("Enter something: ")

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("-> ")
	// 	input, _ := reader.ReadString('\n')
	// 	ch <- input
	// 	time.Sleep(time.Second)
	// }

	// Nested Loops
	for i := 1; i <= 10; i++ {
		fmt.Print("i is ", i)
		for j := 1; j <= 3; j++ {
			fmt.Print("   j: ", j)
		}
		fmt.Println()
	}
}
