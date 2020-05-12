package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Animal struct {
	Name string
	Age  int
}

func (a Animal) String() string {
	return fmt.Sprintf("%v :%d", a.Name, a.Age)
}

type Person struct {
	Name    string
	isExist bool
}

func (p Person) String() string {
	if p.isExist != true {
		return fmt.Sprintf("The person's name is %v.\nthe %v is not alive ", p.Name, p.Name)
	}
	return fmt.Sprintf("The person's name is %v.\n%v is alive", p.Name, p.Name)
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) String() string {
	return fmt.Sprintf("The length is %v.\nThe breadth is %v", r.length, r.breadth)
}

func main() {
	a := Animal{
		Name: "hoge",
		Age:  10,
	}
	fmt.Println(a)

	p := Person{
		Name:    "Gopher",
		isExist: true,
	}
	fmt.Println(p)

	r := Rectangle{
		length:  100,
		breadth: 50,
	}
	fmt.Println(r)

}
