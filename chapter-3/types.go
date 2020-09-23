package types

import "fmt"

func main() {
	fmt.Println("1+1 =", 1+1)            // Sum of 2 integers.
	fmt.Println('H')                     // Prints ascii code of H.
	fmt.Println(len("hey there golang")) // Prints length of the string.
	fmt.Println("str1" + "str2")         // Connects two strings together. *(not addition)*
	fmt.Println("Hello World"[1])        //	The second character of the string Hello World ascii code.
	//Boolean logical operators :
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(true && false)
	fmt.Println(true && true)
	fmt.Println(false && false)
	fmt.Println(!false)
}
