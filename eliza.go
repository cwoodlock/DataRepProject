//Data Representation Project
//Colm Woodlock
//G00341460

package main

//imports for Project
import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/ajax", recieveAjax)

	http.ListenAndServe(":8080, nil")

}
