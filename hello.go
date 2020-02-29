package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", viewHandler)

	http.HandleFunc("/edit/", editHandler)

	//http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return ioutil.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, error := ioutil.ReadFile(fileName)
	if error != nil {
		return nil, error
	}
	return &Page{Title: title, Body: body}, nil
}

func editHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(writer, p)

	// fmt.Fprintf(writer, "<h1>Editing %s</h1>"+
	// 	"<form action = \"/save/%s\" method =\"POST\">"+
	// 	"<textarea namne==\"body\"> %s</textarea><br>"+
	// 	"<input type = \"submit\" value=\"Save\">"+
	// 	"</form>",
	// 	p.Title, p.Title, p.Body)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, _ := loadPage(title)

	t, _ := template.ParseFiles("view.html")
	t.Execute(w, page)
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func handler(responceWriter http.ResponseWriter, request *http.Request) {
	fmt.Println(responceWriter, "Hi there, I love %s!", request.URL.Path[1:])
}
