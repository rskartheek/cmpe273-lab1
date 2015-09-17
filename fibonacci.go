package main

func fibonacci(x uint) uint {
	if x <= 2 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
