package main

import "fmt"

func main() {
	var name string
	var token string

	fmt.Println("Введите ваше имя")
	fmt.Scan(&name)

	fmt.Println("Крестик или нолик?Введите X или O.")
	fmt.Scan(&token)

	fmt.Println(name, token)
}
