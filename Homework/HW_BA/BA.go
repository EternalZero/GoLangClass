/******************************************************************************
 Homework BA
 Create a webpage that serves a form and allows the user to upload a txt file. 
 You do not need to check if the file is a txt; bad programming but just trust 
 the user to follow the instructions. Once a user has uploaded a txt file, copy 
 the text from the file and display it on the webpage. Use req.FormFile and 
 io.Copy to do this.
 Jai
*******************************************************************************/

package main
import (
		"net/http"
		"io"
		"fmt"
)

func textFile(res http.ResponseWriter,req *http.Request){
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res,`
		<form method = "POST"  enctype="multipart/form-data">
		<input type="file" name="name"><br><br>
		<input type="submit">
		</form>`)
	if req.Method == "POST"{
		key := "name"
		_, header, err := req.FormFile(key)
		if err != nil{
			fmt.Println(err)
		}
		inFile, err := header.Open()
		if err != nil{
			fmt.Println(err)
		}
		io.Copy(res,inFile)		
	}
}

func main(){
	http.HandleFunc("/",textFile)
	http.ListenAndServe(":8080",nil)
}