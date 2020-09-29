package main

import (
	"fmt"
	"time"
)

type funcarray func([]int) []int

func swap(a int, b int) {
	var c int
	c = b
	b = a
	a = c
}

func bubblesort(array []int) []int {
	for i := 0; i < len(array); i++ {
		if array[i] < array[i+1] {
			swap(array[i], array[i+1])
		}
	}
	return array
}

func testFunc(a funcarray) funcarray {
	return func(input []int) []int {
		now := time.Now()
		before := now.Nanosecond()
		tmp := a(input)
		after := now.Nanosecond()
		fmt.Println(after - before)
		return tmp
	}
}
