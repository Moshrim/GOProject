package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("MyString is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call MyString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
