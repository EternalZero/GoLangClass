/******************************************************************************
 PROJECT STEP 6 - create a page which illustrates what happens when a user 
 changes a cookie. You can hard-code a changed cookie into your code.
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
	"crypto/hmac"
	"crypto/sha256" 
	"strings"
	"io"
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
				Value: id.String() + "|" + getCode(id.String()) + "|"  + encode + "|" + getCode(encode),
				//Secure: true,
				HttpOnly: true,
				}
			http.SetCookie(res, cookie)
		}
		
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		cv := strings.Split(cookie.Value, "|")
		orgUuid := cv[0]
		orgCodeUuid := cv[1]
		if (getCode(orgUuid) == orgCodeUuid) {
			fmt.Fprintf(res, "The uuid has not been tampered with.\n")
		} else {
			fmt.Fprintf(res,"The uuid has been changed.\n")
		}

		// User changes uuid.
		temp := cv[0]+"123"
		for i:=1; i < len(cv); i++ {
			temp = temp + "|" + cv[i]
		} 
		cookie.Value = temp
		
		// cookie would be reset at this time (but is not done here).
						
		dv := strings.Split(cookie.Value, "|")
		newUuid := dv[0]
		newCodeUuid := dv[1]
		if (getCode(newUuid) == newCodeUuid) {
			fmt.Fprintf(res, "The uuid has not been tampered with.\n")
		} else {
			fmt.Fprintf(res,"The uuid has been changed.\n")
		}
			
		err = tpl.Execute(res,nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	
	http.ListenAndServe(":8080", nil)

}	

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("raka02asd2j01djmk"))
	io.WriteString(h,data)
	return fmt.Sprintf("%x", h.Sum(nil))
	}
	
