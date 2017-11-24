//Data Representation Project
//Colm Woodlock
//G00341460

package main

//imports for Project
import (
	"fmt"
	"html/template"
	"net/http"
)

//Main function
func main() {

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/index", requestHandler)
	http.HandleFunc("/ajax", recieveAjax)

	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("chat.html")
	t.Execute(w, t)
}

//Adapted from https://www.socketloop.com/tutorials/golang-jquery-ajax-post-data-to-server-and-send-data-back-to-client-example
func recieveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ajaxData := r.FormValue("data") //Get the data from the input field in the chat html
		fmt.Fprint(w, elizaResponce(ajaxData))
	}
}

func elizaResponce(userResponce string) string {
	return "hello"
}
