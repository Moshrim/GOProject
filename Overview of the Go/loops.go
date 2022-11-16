package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Аполлон", "Юпитерский", "123@example.com", 23})
	users = append(users, User{"Зевс", "Юпитерский", "321@example.com", 1000})
	users = append(users, User{"Артемида", "Юпитерская", "231@example.com", 18})
	users = append(users, User{"Аид", "Юпитерский", "666@example.com", 561})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

	/*var firstLine = "Once upon a midnight dreary"
	//В Go String -  slice of bytes, поэтому в выводе не буквы//
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	/*animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"

	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}

	/*animals := []string{"dog", "fish", "horse", "cow", "cat"}

	 for _, animal := range animals {
		log.Println(animal)
	} */

	/* for i, animal := range animals {
		log.Println(i, animal)
	} */
}
