package main

import "fmt"

func arrays() {
	var index [5]int
	index2 := [5]int{1, 2, 3, 4, 5}
	//2 ways of defining arrays
	for i := 0; i < 5; i++ {
		index[i] = 6
	}
	fmt.Println(index)
	fmt.Println(index2)

}

func arrayInput() {
	var index [5]int
	var total float64 = 0
	for i := 0; i < 5; i++ {
		fmt.Println("Enter", i, "'s index :")
		fmt.Scanf("%d", &index[i])
		total += float64(index[i])
	}
	fmt.Println("Your array is", index)
	fmt.Println("The average of array is :", (total / float64(len(index))))

	for _, value := range index {
		total += float64(value)
	}
	fmt.Println(total / float64(len(index)))
}
