package main

import "fmt"

func main() {
	var name string
	var token string

	fmt.Println("Введите ваше имя")
	fmt.Scanln(&name)

	fmt.Println("Крестик или нолик?Введите X или O.")
	fmt.Scanln(&token)

	fmt.Println(name, token)
}
