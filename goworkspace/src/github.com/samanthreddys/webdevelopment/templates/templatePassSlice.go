package main

import (
	"html/template"
	"log"
	"os"
	"strconv"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tpl2.gohtml"))
}

var slice1 []string

func main() {

	for i := 0; i < 10; i++ {
		//fmt.Println("String:", i, strconv.Itoa(i))
		slice1 = append(slice1, strconv.Itoa(i))
	}
	err := tmpl.Execute(os.Stdout, slice1)
	if err != nil {
		log.Fatalln(err)
	}

}
