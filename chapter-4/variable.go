package main

import "fmt"

func funcprint() {
	fmt.Println("Function Print")
}

func insert() {
	var input float64
	fmt.Print("Enter a number : \n")
	fmt.Scanf("%f", &input)
	fmt.Println(input * 2)
}

func main() {
	var (
		a = 10
		b = "str"
		c = 0.5
	)

	var age int = 14
	name := "Ted"
	fmt.Println(name, "is", age, "years old")
	var quote string = "Hello World"
	fmt.Println(quote)
	funcprint()
	fmt.Println(a, b, c)
}
