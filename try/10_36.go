package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

var tmpl = template.Must(template.New("msg").Parse("<html><body>{{.Name}}さんの運勢は「<b>{{.Fortune}}</b>」です</body></html>"))

type Result struct {
	Name string
	Fortune string
}

func handler(w http.ResponseWriter, r *http.Request) {
	result := Result{
		Name:r.FormValue("p"),
		Fortune: fortune(),
	}
	tmpl.Execute(w,result)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
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
