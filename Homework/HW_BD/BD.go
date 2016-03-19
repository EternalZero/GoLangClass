/******************************************************************************
 Homework BD
 Create a webpage which writes a cookie to the client's machine. Though this 
 is NOT A BEST PRACTICE, you will store some session data in the cookie. Make 
 sure you use HMAC to ensure that session data is not changed by a user.
 Jai
*******************************************************************************/

package main
import (
		"net/http"
		"html/template"
		"fmt"
		"crypto/hmac"
		"crypto/sha256"
		"io/ioutil"
		"log"
)

var hTemplate string

//load html template
func init() {
	hfile, _ := ioutil.ReadFile("form.gohtml")
	hTemplate = string(hfile)
}
	
func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("raka02asd2j01djmk"))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func session(res http.ResponseWriter,req *http.Request){
	cookie, err := req.Cookie("Session")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "Session",
			Value: " ",
			HttpOnly: true,
		}
	}
	if req.FormValue("music") != "" {
		addInfo := req.FormValue("music")
		cookie.Value = addInfo + " | " + getCode(addInfo)
	}
	http.SetCookie(res, cookie)
	
	fmt.Print(cookie) 
	
	tpl, _ := template.New("hform").Parse(hTemplate)
	err = tpl.Execute(res,nil)
	if err != nil {
		log.Fatalln(err)
	}
	
}

func main(){
	http.HandleFunc("/",session)
	http.ListenAndServe(":8080",nil)
}