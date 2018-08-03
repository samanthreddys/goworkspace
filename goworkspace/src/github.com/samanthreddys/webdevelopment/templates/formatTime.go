package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("tpl5.gohtml").Funcs(fm).ParseFiles("tpl5.gohtml"))

}

var fm = template.FuncMap{
	"FDateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("2006-01-02")

}
func main() {

	err := tmpl.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
