
package main

import (
	"fmt"
	"time"

	"github.com/tenntenn/greeting"
	greeting2 "github.com/tenntenn/greeting/v2"
)

func main() {
	var greet string
	greet = greeting.Do()
	fmt.Printf("%v\n", greet)

	var greet2 string
	greet2 = greeting2.Do(time.Now())
	fmt.Printf("%v\n", greet2)
}
