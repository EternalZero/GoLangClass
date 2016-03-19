/*****************************************************************
 Homework AW:
 Create a webpage that displays the URL path using req.URL.Path
 Jai
******************************************************************/

package main
import (
		"net/http"
		"fmt"
)

func balony(res http.ResponseWriter,req *http.Request){
	fmt.Fprint(res,req.URL.Path)
}

func main(){
	http.HandleFunc("/",balony)
	http.ListenAndServe(":9000",nil)
}