package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

type car struct {
	Model string
	Make  string
}

var tmpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": FirstThree,
}

// FirstThree Function
func FirstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[0:]
	return s

}

func init() {
	tmpl = template.Must(template.New("tpl4.gohtml").Funcs(fm).ParseFiles("tpl4.gohtml"))
}

func main() {
	b := car{
		Model: "Accord",
		Make:  "Honda",
	}
	g := car{
		Model: "Cooper",
		Make:  "Mini",
	}
	l := car{
		Model: "Camry",
		Make:  "Toyota",
	}
	vehicles := []car{b, g, l}
	data := struct {
		Cars []car
	}{
		vehicles,
	}

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
