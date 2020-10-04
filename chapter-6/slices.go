package main

import "fmt"

func slices() {
	var input int
	var temp float64
	var example []float64
	fmt.Println("How large is your array?")
	fmt.Scanf("%d", &input)
	for i := 0; i <= input; i++ {
		fmt.Println("Enter", i, "'s element.")
		fmt.Scanf("%f", &temp)
		/*example = make([]float64, input)
		if i == 0 {
			fmt.Println(example)
		}*/
		example = append(example, temp)
	}
	fmt.Println(example)
}
