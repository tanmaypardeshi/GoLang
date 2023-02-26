package main

import "fmt"

// not allowed
// token := "abcdef"

// allowed
var token string = "abcdef"

// Constant
const PI float32 = 3.14

func main() {
	var username string = "tanmaypardeshi"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T\n", username)

	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("Variable is of type %T\n", isTrue)

	var smallValue uint8 = 8
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type %T\n", smallValue)

	// default values
	var defaultValue int
	fmt.Println(defaultValue)
	fmt.Printf("Variable is of type %T\n", defaultValue)

	// implicit type
	var sampleString = "GoLang"
	fmt.Println(sampleString)
	fmt.Printf("Variable is of type %T\n", sampleString)

	// no var style - This does not work outside of a function
	randomNumber := 30000
	fmt.Println(randomNumber)
	fmt.Printf("Variable is of type %T\n", sampleString)

	// access constant
	fmt.Println(PI)

}
