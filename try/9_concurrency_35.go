package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(r io.Reader) <-chan string {
	ch1 := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch1 <- s.Text()
		}
		close(ch1)
	}()
	return ch1
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Printf(">> %v\n", <-ch)
	}
}
