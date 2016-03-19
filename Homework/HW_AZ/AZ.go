/********************************************************************************
 Homework AZ
 Create a webpage that serves a form and allows the user to enter their name. 
 Once a user has entered their name, show their name on the webpage. Use 
 req.FormValue to do this
 Jai
*********************************************************************************/

package main

import (
	"net/http"
	"io"

)

func prompt(res http.ResponseWriter, req *http.Request) {
	key := "n"
	val := req.FormValue(key)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<form method = "POST">
		<input type = "text" name="n">
		<input type = "submit">
		</form>`)
	io.WriteString(res,`<h1>` + val + `</h1><br>`)
}

func main() {
	http.HandleFunc("/", prompt)
	http.ListenAndServe(":8080", nil)
}