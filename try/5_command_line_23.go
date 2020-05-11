package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
