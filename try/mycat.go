package main

import (
    "fmt"
		"os"
		"io/ioutil"
)

func useIoutilReadFile(fileName string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
			panic(err)
	}
	fmt.Println(string(bytes))
}

func main() {
		args := os.Args
		fileArray := args[1:]
		for _,fn := range fileArray {
		 	useIoutilReadFile(fn)
		}
}