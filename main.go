package main

import (
	"fmt"
	"time"
)

func main() {
	//Fibonacci Series
	fmt.Println(fibonacci(6))
	//Perimeter
	r := Rect{x: 5, y: 5}
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
	fmt.Println(totalPerimeter(&c, &r))
	//Sleep
	sleep(2 * time.Second)
	fmt.Println("Slept 2 seconds")
}
