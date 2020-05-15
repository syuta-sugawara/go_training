package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fortune())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func fortune() (result string) {
	m := map[int]string{
		0: "大吉",
		1: "中吉",
		2: "小吉",
		3: "末吉",
		4: "吉",
		5: "凶",
		6: "大凶",
	}
	rand.Seed(time.Now().UnixNano())
	fortuneNum := rand.Intn(7)
	result = m[fortuneNum]
	return
}
