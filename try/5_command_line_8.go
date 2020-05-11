package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	fmt.Printf("%T\n", argsWithProg)
	fmt.Printf("%T\n", argsWithProg[0])

}
