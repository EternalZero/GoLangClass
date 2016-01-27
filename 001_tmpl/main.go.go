package main

import (
	"log"
	"os"
	"text/template"
)

type shoplist struct {
	Item string
	Num  int
}

type shopcart struct {
	shoplist
	StillNeed bool
}

func main() {
	p1 := shopcart{
		shoplist: shoplist{
			Item: "Tomatoes",
			Num:  3,
		},
		StillNeed: true,
	}

	templ, err := template.ParseFiles("templ.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = templ.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
