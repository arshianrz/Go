package main

// FuncInt is a shitty type
type FuncInt func(int) int

func memorized(fn FuncInt) FuncInt {
	cache := make(map[int]int)

	return func(input int) int {
		if val, found := cache[input]; found {
			println("Read from cache")
			return val
		}

		tmp := fn(input)
		cache[input] = tmp
		return tmp
	}

}
