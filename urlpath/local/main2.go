package main
import (
	"fmt"
	"net/http"
	"io"
)


func Snoop (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<h1>`+req.URL.Path+`</h1><br>`)
}


func main() {


	http.HandleFunc("/",func (res http.ResponseWriter,req *http.Request ) {
		key := "n"
		val := req.FormValue(key)
		fmt.Println("value: ", val)
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		io.WriteString(res, `<h1>`+val+`</h1><br>`)
	})

	http.ListenAndServe(":8080", nil)
}

