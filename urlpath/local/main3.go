package main
//webpage name display
import(
	"fmt"
	"net/http"
	"io"
)

func main() {


	http.HandleFunc("/",func (res http.ResponseWriter,req *http.Request ) {
		key := "n"
		val := req.FormValue(key)
		fmt.Println("value: ", val)
		res.Header().Set("Content-type", "text/html; charset=utf-8")
		io.WriteString(res, `<form
			method = "GET">

			<input type="text" name="n">
			<input type="submit">

			</form>`)

		io.WriteString(res, `<h1>`+val+`</h1><br>`)
	})

	http.ListenAndServe(":8080", nil)
}
