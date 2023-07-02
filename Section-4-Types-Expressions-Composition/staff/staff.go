package staff

import "log"

var OverPaidLimit int = 50000
var underPaidLimit = 10000
type Employee struct {
	FirstName string
	LastName string
	Salary int
	FullTime bool
}

type Office struct {
	AllStaff []Employee
}

func (o *Office) All() []Employee {
	return o.AllStaff
}

func (o *Office) OverPaid() []Employee {
	var overPaid []Employee
	for _, x := range o.AllStaff {
		if x.Salary > OverPaidLimit {
			overPaid = append(overPaid, x)
		}
	}
	return overPaid
}

func (o *Office) UnderPaid() []Employee {
	var underPaid []Employee
	myFunction()
	for _, x := range o.AllStaff {
		if x.Salary <= underPaidLimit {
			underPaid = append(underPaid, x)
		}
	}
	return underPaid
}

func (o *Office) notVisible() {
	log.Println("Hello, World!")
}

func myFunction() {
	log.Println("I am a function")
}