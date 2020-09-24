package main

import "fmt"

func evenOdd() {

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

}

func numerical() {
	for i := 0; i <= 5; i++ {
		switch i {
		case 0:
			fmt.Println("zero")
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
		case 4:
			fmt.Println("four")
		case 5:
			fmt.Println("five")
		default:
			fmt.Println("Unknown Number")
		}
	}
}

func main() {
	var input int
jump:
	fmt.Println("Choose 1 or 2")
	fmt.Scanf("%d", &input)
	if input == 1 {
		evenOdd()
	} else if input == 2 {
		numerical()
	} else {
		goto jump
	}
}
