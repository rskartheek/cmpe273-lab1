package main

import "fmt"
import "time"

func sleep(d time.Duration) {
	c := time.After(d)
	fmt.Println(c)
	<-c
}
