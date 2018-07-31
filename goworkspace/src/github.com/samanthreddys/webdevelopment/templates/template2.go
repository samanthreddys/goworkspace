package main

// using template.Must
import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("*"))

}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template1.go", nil)
	err = tpl.Execute(os.Stdout, "Sam")
	if err != nil {
		log.Fatalln(err)
	}
}
