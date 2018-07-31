package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

type cars struct {
	Model string
	Make  string
	Year  int
}

func init() {
	tmpl = template.Must(template.ParseFiles("tpl3.gohtml"))
}

func main() {
	mini := cars{
		Model: "cooper",
		Make:  "mini",
		Year:  2018,
	}

	err := tmpl.Execute(os.Stdout, mini)
	if err != nil {
		log.Fatalln(err)
	}
}
