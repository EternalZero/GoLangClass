/******************************************************************************
 Homework BF
 Create a web page which serves at localhost over https using TLS.
 Jai
*******************************************************************************/

package main
import (
	"net/http"
	"io"
)

func securePage(res http.ResponseWriter,req *http.Request){
	io.WriteString(res, "This is an example of HTTPS/TLS (without a signed certificate.)")
}

func main(){
	go http.ListenAndServe(":8080", http.RedirectHandler("https://127.0.0.1:10443/",301))
	http.HandleFunc("/", securePage)
	http.ListenAndServeTLS(":10443", "server.pem", "server.key", nil)
}

//openssl genrsa -out server.key 2048
//openssl req -new -x509 -key server.key -out server.pem -days 3650