package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Аполлон",
		LastName:  "Юпитерский",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	var NewVAr float32

	NewVar = 11.1

	/*myMap := make(map[string]iPrintln()
	myMap["Первый"] = 1
	myMap["Второй"] = 2
	log.Println(myMap["Первый"], myMap["Второй"])*/

	/*myMap := make(map[string]string)
	myMap["моя собака"] = "Рипли"
	myMap["моя кошка"] = "Буся"
	myMap["моя собака"] = "Тим"
	log.Println(myMap["моя собака"])
	log.Println(myMap["моя кошка"])*/

}
