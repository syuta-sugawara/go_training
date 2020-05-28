package main

import "fmt"

func Example(name string) string {
	if name == "" {
		return "blank name"
	}
	return fmt.Sprintf("Hello %s", name)
}
