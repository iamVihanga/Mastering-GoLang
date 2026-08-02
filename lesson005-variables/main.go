package main

import "fmt"

// Methods of defining variables in golang

var globalVar int = 10

func main() {
	fmt.Println("Global Variable : ", globalVar)

	var age1 int
	fmt.Println("Age 1 : ", age1)

	var age2 int = 10
	fmt.Println("Age 2 : ", age2)

	var age3 = 10
	fmt.Println("Age 3 : ", age3)

	var age4, name1 = 10, "Vihanga"
	fmt.Printf("%d %s \n", age4, name1)

	age5 := 10
	fmt.Println("Age 5: ", age5)

	age6, name2 := 10, "Vihanga"
	fmt.Printf("%d %s \n", age6, name2)

	something()

	// Variable grouping
	var (
		myInt    int    = 10
		myString string = "Hello world !"
	)

	fmt.Println(myInt, myString)
}

func something() {
	fmt.Println("Global variable access from something function : ", globalVar)
}
