package main
//UploadFile

import(
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	"io"
)


func main(){
	http.HandleFunc("/hs", func(res http.ResponseWriter, req *http.Request){
		key:="q"
		file, _, err := req.FormFile(key)
		if err !=nil {
			log.Println("Error with file", err)
			return
		}
	defer file.Close()
	bs, _ := ioutil.ReadAll(file)
	fmt.Println(string(bs))
	io.WriteString(res,string(bs))
	http.Redirect(res, req, "/",303)
	})


	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html;charset=utf-8")
		io.WriteString(res, `<form method="Post"
			enctype="multipart/form-data"
		action="/hs">
		<input type="file" name="q">
		<input type="submit">
	</form>`)
	})

	http.ListenAndServe(":8080", nil)
}
