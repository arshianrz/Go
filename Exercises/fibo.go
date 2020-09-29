package main

var cache = map[int]int{
	0: 0,
	1: 1,
	2: 1,
	3: 2,
}

func fibo(a int) int {
	if value, ok := cache[a]; ok {
		return value
	}
	result := fibo(a-1) + fibo(a-2)
	cache[a] = result
	return result
}
