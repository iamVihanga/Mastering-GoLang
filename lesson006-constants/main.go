package main

import "fmt"

const x = 10
const y int32 = 15

// Constants can be grouped
const (
	p           int    = 340
	q           uint8  = 245
	programName string = "Lesson 06 - Constants"
	isRunning   bool   = false
)

func main() {

	var a int = x
	fmt.Println(a)

	var b float64 = x
	fmt.Println(b)

	var c int
	// c = y - cannot do this because type incompatibility
	c = int(y)
	fmt.Println(c)

}
