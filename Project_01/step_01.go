/*************************************************************************
 PROJECT STEP 1 - create a web application that serves an HTML template.
**************************************************************************/

package main

import (
	"net/http"
	"html/template"
	"log"
)

type person struct {
	Name string
	Age int
	}
	
func main() {

	p1 := person{
			Name: "Jai",
			Age: 21,
		}

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request) {

		err := tpl.Execute(res,p1)
		if err != nil {
			log.Fatalln(err)
		}
	})
	
	http.ListenAndServe(":8080", nil)

}	
