package main

import "fmt"

func arrays() {
	var example [5]int
	index2 := [5]int{1, 2, 3, 4, 5}
	//2 ways of defining arrays
	for i := 0; i < 5; i++ {
		example[i] = 6
	}
	fmt.Println(example)
	fmt.Println(index2)

}

func arrayInput() {
	var example [5]int
	var total float64 = 0
	for i := 0; i < 5; i++ {
		fmt.Println("Enter", i, "'s example :")
		fmt.Scanf("%d", &example[i])
		total += float64(example[i])
	}
	fmt.Println("Your array is", example)
	fmt.Println("The average of array is :", (total / float64(len(example))))

	total = 0
	for _, value := range example {
		total += float64(value)
	}
	fmt.Println(total / float64(len(example)))
}
