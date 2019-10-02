package main

import "fmt"

func main() {
	var locale string = "es"
	var greeting string = "Yo"

	fmt.Print("Please enter a language: ")
	fmt.Scanf("%s", &locale)

	if locale == "en" {
		greeting = "Hello"
	} else if locale == "es" {
		greeting = "Hola"
	} else if locale == "de" {
		greeting = "Gutentag"
	}

	fmt.Printf(greeting + " Go!\n")
}
