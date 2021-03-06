/******************************************************************************
 PROJECT STEP 5 - continuing to build our application, integrate HMAC into our 
 application to ensure that nobody tampers with the cookie. 
 Jai
*******************************************************************************/

package main

import (
	"net/http"
	"html/template"
	"log"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"crypto/hmac"
	"crypto/sha256" 
)

type User struct {
	Name string
	Age string
	}
	
func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("raka02asd2j01djmk"))
	io.WriteString(h,data)
	return fmt.Sprintf("%x", h.Sum(nil))
	}
		
func main() {

	http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request) {
	
		//parse template
		tpl, err := template.ParseFiles("form.gohtml")
		if err != nil {
			log.Fatalln(err)
		}
	
		person := User{
			Name: req.FormValue("name"), 
			Age: req.FormValue("age"),
		}
	
		b, err := json.Marshal(person)
		if err != nil {
			fmt.Printf("error: " , err)}
		encode := base64.StdEncoding.EncodeToString(b)

		// Set up new cookie if necessary
		cookie, err := req.Cookie("session-fino")
		if err != nil {
			id,_ := uuid.NewV4()
			cookie = &http.Cookie{
				Name: "session-fino",
				Value: id.String() + "|" + getCode(id.String()),
				//Secure: true,
				HttpOnly: true,
				}
			http.SetCookie(res, cookie)
		}
		// Update Cookie
		cookie.Value = cookie.Value + "|" + encode + "|" + getCode(encode)
		http.SetCookie(res, cookie)

		err = tpl.Execute(res,nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	
	http.ListenAndServe(":8080", nil)

}	
