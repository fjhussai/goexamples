package main

import "fmt"

func main() {
	fmt.Println("Do you like Pizza? \n")
	var likespizza string
	fmt.Scanln(&likespizza)

	if (likespizza) == "no" {
		fmt.Println("Thank you for your response")
	} else {
		fmt.Println("What kind of Pizza do you like?")
		var kindpizza string
		fmt.Scanln(&kindpizza)
	}

}
