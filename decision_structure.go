package main

import "log"

func main() {

	myVar := "лошадь"

	switch myVar {
	case "кошка":
		log.Println("кошка - это кошка")

	case "собака":
		log.Println("кошка - это собака")

	case "рыбка":
		log.Println("кошка - это рыбка")

	default:
		log.Println("кошка - что-то другое")
	}
}

/* 	myNum := 100
isTrue := false

			/* 	&& - оба условия должны выполнять одновременно
			! перед переменной означает "нет" */

/* if myNum > 99 && !isTrue {
	log.Println("myNum больше 99 и isTrue - правда")
} else if myNum < 100 && isTrue {
	log.Println("1")
} else if myNum == 101 || isTrue {
	log.Println("2")
} else if myNum > 1000 && isTrue == false {
	log.Println("3")
} */

/* cat := "собака"

if cat == "кошка" {
	log.Println("Кошка есть кошка")
} else {
	log.Println("Кошка не кошка")
} */

/* 			var isTrue bool

   			isTrue = false
   			if isTrue == true {
   				log.Println("isTrue is", isTrue)
   			} else {
   				log.Println("isTrue is", isTrue)
   			} */
