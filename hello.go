package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// http.HandleFunc("/view/", viewHandler)

	// http.HandleFunc("/edit/", editHandler)

	http.HandleFunc("/options/", showOptions)

	http.HandleFunc("/showSelectedOption/", showSelectedOption)

	//http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

// type Page struct {
// 	Title string
// 	Body  []byte
// }

func showOptions(writer http.ResponseWriter, request *http.Request) {
	jsonFile, err := os.Open("options.json")
	jsonByteValue, _ := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	type Options struct {
		OptionOne   string
		OptionTwo   string
		OptionThree string
	}

	var optionsarray Options

	json.Unmarshal(jsonByteValue, &optionsarray)

	template, _ := template.ParseFiles("chooseOption.html")

	template.Execute(writer, optionsarray)

}

func showSelectedOption(writer http.ResponseWriter, request *http.Request) {

	type SelectedOptions struct {
		SelectedRadiobutton string
		AdditionalInfo      string
	}

	myOptions := SelectedOptions{SelectedRadiobutton: request.FormValue("chooseHero"),
		AdditionalInfo: request.FormValue("Thename")}

	template, _ := template.ParseFiles("showSelectedOption.html")
	template.Execute(writer, myOptions)
}

// func (p *Page) save() error {
// 	fileName := p.Title + ".txt"
// 	return ioutil.WriteFile(fileName, p.Body, 0600)
// }

// func loadPage(title string) (*Page, error) {
// 	fileName := title + ".txt"
// 	body, error := ioutil.ReadFile(fileName)
// 	if error != nil {
// 		return nil, error
// 	}
// 	return &Page{Title: title, Body: body}, nil
// }

// func editHandler(writer http.ResponseWriter, request *http.Request) {
// 	title := request.URL.Path[len("/edit/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		p = &Page{Title: title}
// 	}
// 	t, _ := template.ParseFiles("edit.html")
// 	t.Execute(writer, p)

// }

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	page, _ := loadPage(title)

// 	t, _ := template.ParseFiles("view.html")
// 	t.Execute(w, page)
// }

// func handler(responceWriter http.ResponseWriter, request *http.Request) {
// 	fmt.Println(responceWriter, "Hi there, I love %s!", request.URL.Path[1:])
// }
