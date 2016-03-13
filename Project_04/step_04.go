/******************************************************************************
 PROJECT STEP 4 - refactoring our application, create a new data type called 
 "user" which has fields for the user's name and age. When you receive the 
 user's name and age form submission, create a variable of type "user" then 
 put those values from the form submission into the fields for that variable. 
 Marshal your variable of type "user"  to JSON. Encode that JSON to base64. 
 Store that value in the cookie.
 Jairo Reyes
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
)

type User struct {
	Name string
	Age string
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
				Value: id.String(),
				//Secure: true,
				HttpOnly: true,
				}
			http.SetCookie(res, cookie)
		}
		// Update Cookie
		cookie.Value = cookie.Value + "|" + encode
		http.SetCookie(res, cookie)

		err = tpl.Execute(res,nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	
	http.ListenAndServe(":8080", nil)

}	
