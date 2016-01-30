package main

import(

	"net/http"
	"html/template"
	"log"
)

func surfsup(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res,nil)
	}




func main() {

	http.HandleFunc("/", surfsup)

	http.Handle("/surfing/", http.StripPrefix("/surfing", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}


