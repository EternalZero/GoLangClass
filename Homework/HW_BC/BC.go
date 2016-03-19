/******************************************************************************
 Homework BC
 Create a webpage which writes a cookie to the client's machine. This cookie 
 should be designed to create a session and should use a UUID, HttpOnly, and 
 Secure (though you'll need to comment secure out).
 Jai
*******************************************************************************/

package main
import (
		"net/http"
		"fmt"
		"github.com/nu7hatch/gouuid"
)


func createSession(res http.ResponseWriter,req *http.Request){
	cookie, err := req.Cookie("session")
	if err != nil {
		id,_ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res,cookie)
	}
	fmt.Fprint(res,cookie)
}

func main(){
	http.HandleFunc("/",createSession)
	http.ListenAndServe(":8080",nil)
}