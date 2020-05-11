package main

import (
	"flag"
	"fmt"
)

var msg = flag.String("msg", "デフォルト値", "説明")

func main() {
	var str string
	flag.StringVar(&str, "string", "初期値", "説明")
	var (
		i = flag.Int("int", 0, "int flag")
		s = flag.String("str", "default", "string flag")
		b = flag.Bool("bool", false, "bool flag")
	)
	flag.Parse()
	fmt.Println(*i, *s, *b, *msg)
	fmt.Println(str)
}
