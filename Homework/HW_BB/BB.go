/******************************************************************************
 Homework BB
 Create a webpage which uses a cookie to track the number of visits of a user. 
 Display the number of visits. Make sure that the favicon.ico requests are not 
 also incrementing the number of visits.
 Jai
*******************************************************************************/

package main
import (
		"net/http"
		"strconv"
		"time"
		"io"
)


func numberOfVisits(res http.ResponseWriter,req *http.Request){
	cookie, err := req.Cookie("visits")
	if err == http.ErrNoCookie{
		cookie = &http.Cookie{Name:"visits", Value:"0"}
	}
	numberVisits, err := strconv.Atoi(cookie.Value)
	numberVisits++
	cookie.Value = strconv.Itoa(numberVisits)
	cookie.Expires = time.Now().Add(500*24*time.Hour)
	http.SetCookie(res, cookie)
	io.WriteString(res, "This site has been visited ")
	io.WriteString(res, cookie.Value)
	io.WriteString(res, " times.")

}

func main(){
	http.HandleFunc("/", numberOfVisits)
	http.HandleFunc("/favicon.ico",func(res http.ResponseWriter,req *http.Request){})
	http.ListenAndServe(":8080",nil)
}