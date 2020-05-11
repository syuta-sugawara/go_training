package main

import (
	"fmt"

	"github.com/tenntenn/greeting"
)

func main() {
	var greet string
	greet = greeting.Do()
	fmt.Printf("%v\n", greet)
}
