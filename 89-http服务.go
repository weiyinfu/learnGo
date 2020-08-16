package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

var editTemplate = `
<h1>Editing {{.Title}}</h1>

<form action="/save/{{.Title}}" method="POST">
<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
`

//运行此程序，打开浏览器，访问localhost:8080/haha
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, request.URL.Path)
	})
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1>hello</h1>")
	})
	http.HandleFunc("/edit", func(writer http.ResponseWriter, request *http.Request) {
		t := template.New(editTemplate)
		fmt.Println("ok")
		t.Execute(writer, Page{Title: "this is title", Body: []byte("this is content")})
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
