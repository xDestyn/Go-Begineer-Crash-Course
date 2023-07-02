package main

import (
	"Section-4-Types-Expressions-Composition/staff"
	"fmt"
	"log"
	"sort"
)

// basic types (numbers, strings, booleans)
// Depending on the architecture of the machine, the size of the int type can be 16, 32 or 64 bits.
var myInt int
var myInt16 int16
var myInt32 int32
var myInt64 int64
var myUint uint

// Float 32 and 64 are the two types of floating-point numbers in Go.
var myFloat32 float32
var myFloat64 float64

// aggregate types (arrays, structs)
// type Car struct {
// 	NumberOfTires int
// 	Luxury        bool
// 	BucketSeats   bool
// 	Make          string
// 	Model         string
// 	Year          int
// }

// reference types (pointers, slices, maps, functions, channels)
// type Animal struct {
// 	Name         string
// 	Sound        string
// 	NumberOfLegs int
// }

// var keyPressChan chan rune

// interface types
type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

// type Employee struct {
// 	Name     string
// 	Age      int
// 	Salary   int
// 	FullTime bool
// }

// Composition
type Vehicle struct {
	NumberOfWheels int
	NumberOfPassengers int
}

type Car struct {
	Make string
	Model string
	Year int
	IsElectric bool
	IsHybrid bool
	Vehicle Vehicle
}

func (v Vehicle) showDetails() {
	log.Println("Number of passengers:", v.NumberOfPassengers)
	log.Println("Number of wheels:", v.NumberOfWheels)
}

func (c Car) show() {
	log.Println("Make:", c.Make)
	log.Println("Model:", c.Model)
	log.Println("Year:", c.Year)
	log.Println("Is electric:", c.IsElectric)
	log.Println("Is hybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}

// Exported vs Unexported
var employees = []staff.Employee {
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 60000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 50000, FullTime: false},
}

func main() {
	// Basic types
	myInt = 10
	myUint = 20

	myFloat32 = 10.1
	myFloat64 = 20.2

	log.Println(myInt, myUint, myFloat32, myFloat64)

	myString := ""

	log.Println(myString)

	myString = "John"

	var myBool = true
	myBool = false

	log.Println(myBool)

	// Aggregate types - Arrays
	var myStrings [3]string
	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	log.Println("First element in array is: ", myStrings[0])
	// Aggregate types - Structs
	// First way to declare a struct.
	// var myCar Car
	// myCar.NumberOfTires = 4
	// myCar.Luxury = false
	// myCar.Make = "Infiniti"

	// Second way to declare a struct.
	// myCar := Car{
	// 	NumberOfTires: 4,
	// 	Luxury:        false,
	// 	BucketSeats:   true,
	// 	Make:          "Infiniti",
	// 	Model:         "Q50",
	// 	Year:          2016,
	// }

	// log.Printf("My car is a %d %s %s.", myCar.Year, myCar.Make, myCar.Model)

	// References types - Pointers
	myInt := 10

	myFirstPointer := &myInt

	log.Println("myInt is", myInt)
	log.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 14
	log.Println("myInt is", myInt)
	log.Println("myFirstPointer is", myFirstPointer)

	changeValueOfPointer(&myInt)
	log.Println("After function call, myInt is now", myInt)

	// References types - Slices
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	log.Println(animals)

	for i, x := range animals {
		log.Println(i, x)
	}

	log.Println("Element 0 is", animals[0])

	log.Println("First two elements are", animals[0:2])

	log.Println("The slice is", len(animals), "elements long")

	log.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)

	log.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	log.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	log.Println(animals)

	// References types - Maps
	// Map is reference type. Always passed by reference.
	// Don't have to use pointers with maps.
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		log.Println(key, value)
	}

	delete(intMap, "four")

	for key, value := range intMap {
		log.Println(key, value)
	}

	el, ok := intMap["four"]
	if ok {
		log.Println("Found element", el)
	} else {
		log.Println("Did not find element")
	}

	intMap["two"] = 4
	log.Println(intMap)

	// Reference Types - Functions
	z := addTwoNumbers(1, 2)
	log.Println(z)

	myTotal := sumMany(1, 2, 3, 4, 5)
	log.Println(myTotal)

	// var dog Animal
	// dog.Name = "Buddy"
	// dog.Sound = "Woof"
	// dog.NumberOfLegs = 4
	// dog.Says()

	// cat := Animal{
	// 	Name:         "Cat",
	// 	Sound:        "Meow",
	// 	NumberOfLegs: 4,
	// }
	// cat.Says()
	// cat.HowManyLegs()

	// Reference Types - Channels
	// go doSomething("Hello, world")

	// log.Println("This is another message")
	// for {
	// 	// Do nothing
	// }
	// keyPressChan = make(chan rune)
	// go listenForKeyPress()

	// Uncomment for channels
	// log.Println("Press any key, or q to quit.")
	// _ = keyboard.Open()
	// defer func() {
	// 	keyboard.Close()
	// }()

	// for {
	// 	char, _, _ := keyboard.GetSingleKey()
	// 	if char == 'q' || char == 'Q' {
	// 		break
	// 	}
	// 	keyPressChan <- char
	// }

	// Reference Types - Interfaces
	// ask a riddle
	dog := Dog{
		Name:         "Buddy",
		Sound:        "Woof",
		NumberOfLegs: 4,
	}
	Riddle(&dog)

	var cat Cat
	cat.Name = "Cat"
	cat.NumberOfLegs = 4
	cat.Sound = "Meow"
	cat.HasTail = true

	Riddle(&cat)

	// Expressions
	age := 10
	name := "Jack"
	rightHanded := true

	rightHanded = false

	log.Printf("%s is %d years old and is right handed: %t", name, age, rightHanded)
	ageInTenYears := age + 10
	log.Printf("In ten years, %s will be %d years old.", name, ageInTenYears)
	isATeenager := age >= 13 && age <= 19
	log.Println(name, "is a teenager:", isATeenager)

	// Booleans
	apples := 18
	oranges := 9

	// boolean expression
	log.Println(apples == oranges)
	log.Println(apples != oranges)

	// > < >= <=
	log.Printf("%d > %d: %t", apples, oranges, apples > oranges)
	log.Printf("%d < %d: %t", apples, oranges, apples < oranges)
	log.Printf("%d >= %d: %t", apples, oranges, apples >= oranges)
	log.Printf("%d <= %d: %t", apples, oranges, apples <= oranges)

	// Compound Boolean Expressions
	// jack := Employee {
	// 	Name: "Jack",
	// 	Age: 25,
	// 	Salary: 50000,
	// 	FullTime: true,
	// }

	// jill := Employee {
	// 	Name: "Jill",
	// 	Age: 35,
	// 	Salary: 80000,
	// 	FullTime: true,
	// }

	// var employees []Employee
	// employees = append(employees, jack)
	// employees = append(employees, jill)

	// for _, x := range employees {
	// 	if x.Age > 30 {
	// 		log.Println(x.Name, "is older than 30")
	// 	} else {
	// 		log.Println(x.Name, "is not older than 30")
	// 	}

	// 	if x.Age > 30 && x.Salary > 60000 {
	// 		log.Println(x.Name, "makes more than 60k and is older than 30")
	// 	} else {
	// 		log.Println("Either", x.Name, "makes less than 60k or is not older than 30")
	// 	}

	// 	if x.Age > 30 || x.Salary > 60000 {
	// 		log.Println(x.Name, "is older than 30 or makes more than 60k")
	// 	} else {
	// 		log.Println("Either", x.Name, "makes less than 60k or is not older than 30")
	// 	}

	// 	if (x.Age > 30 || x.Salary > 60000) && x.FullTime {
	// 		log.Println(x.Name, "matches our unclear criteria")
	// 	}
	// }

	// Composition
	suv := Vehicle {
		NumberOfWheels: 4,
		NumberOfPassengers: 6,
	}

	volvoXC90 := Car {
		Make: "Volvo",
		Model: "XC90",
		Year: 2020,
		IsElectric: false,
		IsHybrid: true,
		Vehicle: suv,
	}

	volvoXC90.show()

	teslaModelX := Car {
		Make: "Tesla",
		Model: "Model X",
		Year: 2020,
		IsElectric: true,
		IsHybrid: false,
		Vehicle: suv,
	}

	teslaModelX.show()

	// Exported vs Unexported
	myStaff := staff.Office{
		AllStaff: employees,
	}

	log.Println(myStaff.All())
	staff.OverPaidLimit = 30000
	log.Println("Overpaid staff", myStaff.OverPaid())
	log.Println("Underpaid staff", myStaff.UnderPaid())

	// myStaff.notVisible()

}

func changeValueOfPointer(num *int) {
	*num = 77
}

func DeleteFromSlice(a []string, i int) []string {
	// Copy first element to index i. ["cat", "dog", "fish"]
	a[i] = a[len(a)-1]
	// Erase last element (write zero value).
	a[len(a)-1] = ""
	// Delete last element (truncate slice).
	a = a[:len(a)-1]
	return a
}

func addTwoNumbers(x, y int) (sum int) {
	// Naked return
	sum = x + y
	return
}

// Variatic Function
func sumMany(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Receiver Function
// func (a *Animal) Says() {
// 	log.Printf("A %s says %s", a.Name, a.Sound)
// }

// func (a *Animal) HowManyLegs() {
// 	log.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
// }

// func doSomething(s string) {
// 	until := 0
// 	for {
// 		log.Println("s is", s)
// 		until = until + 1
// 		if until >= 5 {
// 			break
// 		}
// 	}
// }

// func listenForKeyPress() {
// 	for {
// 		key := <-keyPressChan
// 		log.Println("You pressed", string(key), "key")
// 	}
// }

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it?`, a.Says(), a.HowManyLegs())
	log.Println(riddle)
}
