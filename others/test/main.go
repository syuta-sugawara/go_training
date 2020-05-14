package main

import (
	"fmt"

	"./hello"
)

func main() {
	s := hello.GetHello("hogehoge")
	fmt.Println(s)
}
