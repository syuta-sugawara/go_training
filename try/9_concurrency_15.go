package main

import (
	"fmt"
	"time"
)

func main() {
	var v int
	go func() {
		time.Sleep(1 * time.Second)
		v = 100
	}()
	func() {
		time.Sleep(3 * time.Second)
		fmt.Println(v)
	}()

	time.Sleep(2 * time.Second)
}
