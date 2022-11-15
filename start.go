package main

/*Что=то в файле будет использовать format*/
import "fmt"

func main() {
	fmt.Println("Hello, world!")

	/*Может быть огромное кол-во переменных, нотолько используемых*/
	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world."
	fmt.Println(whatToSay)

	i = 23

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println("Функция вернула", whatWasSaid, theOtherThingThatWasSaid)
}

/*Функция будет что-то возвращать*/
func saySomething() (string, string) {
	return "что-то", "еще"
}
