/****************************************************************************
 PROJECT STEP 3 - continuing to build our application, create a template 
 which is a form. The form should gather the user's name and age. Store the 
 user's name and age in the cookie.
 Jairo Reyes
*****************************************************************************/

package main

import (
	"net/http"
	"html/template"
	"log"
	"github.com/nu7hatch/gouuid"
)

type person struct {
	Name string
	Age string
	}
	
func main() {

	http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request) {
	
		// parse template
		tpl, err := template.ParseFiles("form.gohtml")
		if err != nil {
			log.Fatalln(err)
		}
	
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
		newValue := req.FormValue("name") + " " + req.FormValue("age")
		cookie.Value = cookie.Value + "|" + newValue
		http.SetCookie(res, cookie)

		err = tpl.Execute(res,nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	
	http.ListenAndServe(":8080", nil)

}	
