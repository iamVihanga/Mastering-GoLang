package main

import "fmt"

func main() {
	a := 15
	b := 10
	c := 25.54

	// Sum operator
	sum := a + b
	fmt.Println("Sum : ", sum)

	// Difference operator
	difference := a - b
	fmt.Println("Difference : ", difference)

	// Product operator
	product := a * b
	fmt.Println("Product : ", product)

	// Division operator
	division := a / b
	fmt.Println("Division : ", division)

	// Remainder operator
	remainder := a % b
	fmt.Println("Remainder : ", remainder)

	// Typecasting
	newSum := float64(a) + c
	fmt.Println("New Sum : ", newSum)

	// Computational Shorthand
	a += 5
	b -= 2
	a *= 2
	b %= 2

	fmt.Println(a, b)

	// Division by Zero
	// divisionByZero := a / 0

	// Increment & Decrement Operators
	a++
	b--
	fmt.Println(a, b)
}
