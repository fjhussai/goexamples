package main

import "fmt"

func main() {
	var locale string = "es"
	var greeting string = "Yo"

	/* var languages string = "es" , "en", "de" , "fr" */

	fmt.Scanf("%s", &locale)

	if locale == "en" {
		greeting = "Hello"
	} else if locale == "es" {
		greeting = "Hola"
	} else if locale == "de" {
		greeting = "Gutentag"
	} else if locale == "fr" {
		greeting = "Bonjour"
	}

	fmt.Printf(greeting + " Go!\n")
}
