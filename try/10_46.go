package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080?p=Gopher")
	if err != nil {
		println("error")
	}
	defer resp.Body.Close()
	fmt.Printf("%v\n", resp.Body)
}