package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}

	// shubham := person{firstName: "shubham", lastName: "v"}
	// fmt.Println(shubham)

	// var shubham person
	// shubham.firstName = "shubham"
	// shubham.lastName = "v"
	// fmt.Println(shubham)
	// fmt.Printf("%+v", shubham)

	// shubham := person{
	// 	firstName: "shubham",
	// 	lastName:  "v",
	// 	contactInfo: contactInfo{
	// 		email:   "shubhamvarshney@outlook.com",
	// 		zipCode: 560034,
	// 	},
	// }

	// shubham.updateName("Shubham")

	// shubhamPointer := &shubham
	// fmt.Println(*&shubhamPointer)
	// shubhamPointer.updateName("Shubham")

	// shubham.updateName("Shubham")

	// shubham.print()

	name := "shubham"
	fmt.Println(&name)
	fmt.Println(&(*(&name)))
	namePointer := &name
	fmt.Println(&namePointer)

	printPointer(&name)

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSilce(mySlice)
	fmt.Println(mySlice)

	// Value Types: 	int,    float, string,   bool,     structs
	// Referrence Type: slices, maps,  channels, pointers, functions
}

//pass by referrence, pointer is passed in receiver
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Pass By value, for Receiver, hence no effect in passed person object
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func printPointer(name *string) {
	println(&name)
}

func updateSilce(s []string) {
	s[0] = "Bye"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
