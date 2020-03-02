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

	http.HandleFunc("/options/", showOptions)

	http.HandleFunc("/showSelectedOption/", showSelectedOption)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

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
