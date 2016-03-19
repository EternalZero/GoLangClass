/********************************************************************************
 Homework AY:
 Create a webpage that serves at localhost:8080 and will display the name in the
url when the url is localhost:8080?n="some-name" - use req.FormValue to do this
 Jai
*********************************************************************************/

package main

import (
	"net/http"
	"fmt"

)

func balony(res http.ResponseWriter, req *http.Request) {
	key := "n"
	val := req.FormValue(key)
	fmt.Fprint(res, val)
}

func main() {
	http.HandleFunc("/", balony)
	http.ListenAndServe(":8080", nil)
}