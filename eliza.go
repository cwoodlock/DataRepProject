//Data Representation Project
//Colm Woodlock
//G00341460

package main

//imports for Project
import (
  "fmt"
  "math/rand"
  "time"
  "regex"
  "net/http"
  "text/template"
)

//Main function
func main() {
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.HandleFunc("/", requestHandler)
    http.HandleFunc("/ajax", recieveAjax)

    http.ListenAndServe(":8080, nil")
}

func requestHandler(w http.ResponseWriter, r *http.Request){
  t, _ := requestHandler("chat.html")
  t.Execute(w, t)
}
