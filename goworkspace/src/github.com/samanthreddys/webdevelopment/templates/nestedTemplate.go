package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {

	err := tmpl.ExecuteTemplate(os.Stdout, "tpl6.gohtml", "Tom")
	if err != nil {
		log.Fatalln(err)
	}

}
