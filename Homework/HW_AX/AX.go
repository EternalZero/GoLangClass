/**********************************************************************
 Homework AX:
 Create a webpage that serves at localhost:8080 and will display the 
 name in the url when the url is localhost:8080/name - use req.URL.Path 
 to do this
 Jai
***********************************************************************/

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
	http.ListenAndServe(":8080",nil)
}