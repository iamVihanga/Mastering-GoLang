package main

import "fmt"

func main() {

	// Variable type - uint8
	// Stores positive integers from 0 to 255
	var smallPositiveValue uint8

	smallPositiveValue = 255
	smallPositiveValue = smallPositiveValue + 1 // This will cause an overflow

	fmt.Println("Small positive value:", smallPositiveValue)

	// Variable type - int8
	// Stores integers from -128 to 127
	var smallNegativeValue int8

	smallNegativeValue = -128
	smallNegativeValue = smallNegativeValue - 1 // This will cause an underflow

	fmt.Println("Small negative value:", smallNegativeValue)

	// Variable type - int or uint
	// Depending on the architecture of the machine, int can be either 32 or 64 bits
	var largeValue int

	largeValue = 2147483647     // Maximum value for a 32-bit signed integer
	largeValue = largeValue + 1 // This will cause an overflow on a 32-bit system

	fmt.Println("Large value:", largeValue)

	var largeUnsignedValue uint

	largeUnsignedValue = 4294967295             // Maximum value for a 32-bit unsigned integer
	largeUnsignedValue = largeUnsignedValue + 1 // This will cause an overflow on a 32-bit system

	fmt.Println("Large unsigned value:", largeUnsignedValue)

	// Type-casting integers
	var smallInt int8 = 100
	var largeInt int16

	// Type-casting from int8 to int16
	largeInt = int16(smallInt)

	fmt.Println("Small int:", smallInt)
	fmt.Println("Large int after type-casting:", largeInt)

	// Can also type-cast from int16 to int8, but be careful of overflow
	smallInt = int8(largeInt) // This will work fine since largeInt is within the range of int8

	fmt.Println("Small int after type-casting back:", smallInt)

	// Variable type - Byte
	// Byte is an alias for uint8 and is used to represent a single byte of data
	var byteValue byte = 'A' // ASCII value of 'A' is 65

	fmt.Println("Byte value:", byteValue)
	fmt.Println("Byte value as character:", string(byteValue))

	// Variable type - Rune
	// Rune is an alias for int32 and is used to represent a Unicode code point
	var runeValue rune = '世' // Unicode code point for '世'

	fmt.Println("Rune value:", runeValue)
	fmt.Println("Rune value as character:", string(runeValue))
}
