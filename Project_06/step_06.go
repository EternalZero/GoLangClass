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

		//parse template
		tpl, err := template.ParseFiles("form.gohtml")
		if err != nil {
			log.Fatalln(err)
		}

	http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request) {
		
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
				MaxAge: 3600, 
				//Secure: true,
				HttpOnly: true,
				}
			http.SetCookie(res, cookie)
		}
		// Update Cookie
		//cookie.Value = cookie.Value + "|" + encode + "|" + getCode(encode)
		//http.SetCookie(res, cookie)
		
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		cv := strings.Split(cookie.Value, "|")
		orgUuid := cv[0]
		orgCodeUuid := cv[1]
		fmt.Print("cv[0] = ", cv[0], "\n")
		fmt.Print("cv[1] = ", cv[1], "\n")
		
		if (getCode(orgUuid) == orgCodeUuid) {
			fmt.Fprintf(res, "The uuid has not been tampered with.\n")
		} else {
			fmt.Fprintf(res,"The uuid has been changed.\n")

		}
		
		// User changes uuid and resets cookie.
		temp := cv[0]+"123"
		fmt.Print("temp = ", temp, "\n")
		for i:=1; i < len(cv); i++ {
			temp = temp + "|" + cv[i]
		} 
		fmt.Print("tempFinal = ", temp, "\n")
		cookie.Value = temp
		
		////////////////////////////////////////////////////////////////////////////
		
		fmt.Print("cookie.Value = ", cookie.Value, "\n")  // cookie.Value is CORRECT HERE
		http.SetCookie(res, cookie)  // This does NOT change cookie to have the new cookie.Value
		
		///////////////////////////////////////////////////////////////////////////
		/*
		cookieN, err := req.Cookie("session-fino")
		fmt.Print("cookieN.Value = ", cookieN.Value, "\n")
		dv := strings.Split(cookieN.Value, "|")
		newUuid := dv[0]
		fmt.Print("dv[0] = ", dv[0], "\n")
		newCodeUuid := dv[1]
		fmt.Print("dv[1] = ", dv[1], "\n")
		if (getCode(newUuid) == newCodeUuid) {
			fmt.Fprintf(res, "The uuid has not been tampered with.\n")
		} else {
			fmt.Fprintf(res,"The uuid has been changed.\n")

		}
		*/
			
			

		
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
	
