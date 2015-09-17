package main

import (
	"math"
	"testing"
	"time"
)

//TestFibonacci is an exported func
func TestFibonacci(t *testing.T) {
	var x uint
	x = fibonacci(15)
	if x != 610 {
		t.Error("Expected 610, got ", x)
	}
}

//TestCirclePerimeter is an exported func
func TestCirclePerimeter(t *testing.T) {
	var v float64
	circlePerimeter := Circle{x: 0, y: 0, r: 5}
	v = circlePerimeter.perimeter()
	if v != math.Pi*10 {
		t.Error("Expected 1.5, got ", v)
	}
}

//TestRectPerimeter is an exported func
func TestRectPerimeter(t *testing.T) {
	var x float64
	rectPerimeter := Rect{x: 5, y: 5}
	x = rectPerimeter.perimeter()
	if x != 20 {
		t.Error("Expected 20, got ", x)
	}
}

//TestSleep is an exported test func
func TestSleep(t *testing.T) {
	start := time.Now()
	sleep(2 * time.Second)
	elapsed := time.Since(start)
	if elapsed >= (time.Second)*3 {
		t.Error("Expected 2, got ", elapsed)
	}

}
