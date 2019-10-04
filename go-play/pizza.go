package main

import "fmt"

func main() {
	fmt.Println("Do you like Pizza? \n")
	var answer string //Changed var likespizza to answer, for simplicity.
	fmt.Scanln(&answer)

	if (answer) == "no" {
		fmt.Println("Thank you for your response")
	} else {
		fmt.Println("What kind of Piza do you like?")
	}

}
