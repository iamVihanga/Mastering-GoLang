package main

import "fmt"

/**

- In go lang, strings only can be defined using "" or ``
- zero value of strings in go is an empty string ""
- Go supports Unicode and using UTF-8 for encode unicodes
- UTF-8 encoding means each unicode character can take from 1 to 4 bytes in memory
- rune data type is used to represent a single unicode character with 4 bytes of memory ( int32 )

*/

func main() {
	var myString string

	fmt.Println(myString)

	myString = "Hello World !"

	fmt.Println(myString)

	myString = "Hello\nWorld !"

	fmt.Println(myString)

	myString = `Hello
	
		World !`

	fmt.Println(myString)

	var firstname, lastname string

	firstname = "Vihanga"
	lastname = "Silva"

	var fullname string
	fullname = firstname + " " + lastname

	fmt.Println(fullname)
	fmt.Printf("%s %s \n", firstname, lastname)

	fullname = fmt.Sprintf("%s %s", firstname, lastname)

	fmt.Println(fullname)

}
