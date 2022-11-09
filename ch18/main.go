package main

var cache = map[int]int{}

func Fibonacci(n int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	re := Fibonacci(n-1) + Fibonacci(n-2)
	cache[n] = re
	return re
}
