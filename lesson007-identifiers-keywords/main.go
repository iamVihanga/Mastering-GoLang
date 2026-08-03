package main

import "fmt"

func main() {

	var myInt int = 10
	/**
	- variable name (identifier) cannot start with a digit
	- starting with underscore "_" is valid
	*/

	fmt.Println(myInt)

	var largeNumber int = 10_000_000
	/**
	- When defining large numbers comma shouldnt be used
	- Ex : 10,000,000 is invalid
	- Valid form is using underscore
	*/

	fmt.Println(largeNumber)

	const π float32 = 3.14
	/**
	- Unicode characters are also allowed as identifiers
	*/

	fmt.Println(π)

}
