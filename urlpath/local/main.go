package main
import (
	"net/http"
	"io"
)


func Snoop (res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<h1>`+req.URL.Path+`</h1><br>`)
}


func main() {


	http.HandleFunc("/", Snoop)

	http.ListenAndServe(":8080", nil)
}

