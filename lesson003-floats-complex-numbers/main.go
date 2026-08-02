package main

import "fmt"

/**

Data Type `Float`
 - Floating point numbers have integer value, a decimal point and a fractional value.
 - Unline int data types, float can store fractional values.
 - In go lang, there are two types of `float` such as `float32` and `float64`


Data Type `Complex`
 - Complex numbers in go have two parts, a real part and an imaginary part.
 - `complex64` and `complex128` are the two types of complex numbers in go.
 - `complex64` has a real and imaginary part of type `float32`
 - `complex128` has a real and imaginary part of type `float64`
 - built-in function `complex()` is used to create complex numbers in go.
 - built-in function `real()` is used to extract the real part of a complex number.
 - built-in function `imag()` is used to extract the imaginary part of a complex number.

*/

func main() {

	// Variable type - float32
	// Used to store fractional values with 32 bits of precision
	var smallFloatValue float32

	fmt.Println(smallFloatValue)

	smallFloatValue = 3.1415926535897932384

	fmt.Println("smallFloatValue:", smallFloatValue)

	// Variable type - float64
	// Used to store fractional values with 64 bits of precision
	var largeFloatValue float64

	fmt.Println(largeFloatValue)

	largeFloatValue = 3.1415926535897932384

	fmt.Println("largeFloatValue:", largeFloatValue)

	// Variable type - complex64
	// Used to store complex numbers with 32 bits of precision for both real and imaginary parts
	var smallComplexValue complex64

	fmt.Println(smallComplexValue)

	smallComplexValue = complex(3.1415926535897932384, 2.7182818284590452353)

	fmt.Println("smallComplexValue:", smallComplexValue)
	fmt.Println("Real part:", real(smallComplexValue))
	fmt.Println("Imaginary part:", imag(smallComplexValue))

	// Variable type - complex128
	// Used to store complex numbers with 64 bits of precision for both real and imaginary parts
	var largeComplexValue complex128

	fmt.Println(largeComplexValue)

	largeComplexValue = complex(3.1415926535897932384, 2.7182818284590452353)

	fmt.Println("largeComplexValue:", largeComplexValue)
	fmt.Println("Real part:", real(largeComplexValue))
	fmt.Println("Imaginary part:", imag(largeComplexValue))

}
