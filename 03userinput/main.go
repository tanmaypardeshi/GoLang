package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the a message:- ")

	// comma ok || err ok
	input, error := reader.ReadString('\n')

	fmt.Println("Input is:- ", input)
	fmt.Println("Error is:- ", error)
	fmt.Printf("Type of input is:- %T\n", input)
	fmt.Printf("Type of error is:- %T\n", error)

}
