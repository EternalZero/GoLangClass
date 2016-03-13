/****************************************************************************
 PROJECT STEP 2 - have the application write a cookie called "session-fino" 
 with a UUID. The cookie should serve HttpOnly and you should have the 
 "Secure" flag set also though comment the "Secure" flag out as we're not 
 using https.
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
	Age int
	}
	
func main() {

	p1 := person{
			Name: "Jairo",
			Age: 21,
		}

	// parse template
	tpl, err := template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request) {
	
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

		err = tpl.Execute(res,p1)
		if err != nil {
			log.Fatalln(err)
		}
	})
	
	http.ListenAndServe(":8080", nil)

}	
