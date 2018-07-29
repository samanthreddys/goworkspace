package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	tpl, err := template.ParseFiles("tpl1.gohtml", "template1.go")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)

	err = tpl.ExecuteTemplate(os.Stdout, "template1.go", nil)

	if err != nil {
		log.Fatalln(err)
	}

	// using ParseGlob method
	tpl1, err := template.ParseGlob("*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl1.ExecuteTemplate(os.Stdout, "tpl1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl1.ExecuteTemplate(os.Stdout, "index.html", nil)
}
