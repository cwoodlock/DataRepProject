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

	http.HandleFunc("/", requestHandler)
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
		userResponce := r.FormValue("data") //Get the data from the input field in the chat html
		fmt.Fprint(w, elizaResponce(userResponce))
	}
}

func elizaResponce(userResponce string) string {
	var responces = [][]string{
		{`(?i)^\s*I need ([^.!?]*)[.!?]*\s*$`,
			"Why do you need $1?",
			"Would it really help you to get $1?"},

		{`(?i)^\s*Why don'?t you ([^.!?]*)[.!?\s]*$`,
			"Do you really think I don't $1?",
			"Perhaps eventually I will $1.",
			"Do you really want me to $1?"},

		{`(?i)^\s*Why can'?t I ([^.!?]*)[.!?\s]*$`,
			"Do you think you should be able to $1?",
			"If you could $1, what would you do?",
			"I don't know - why can't you $1?",
			"Have you really tried?"},

		{`(?i)^\s*(?:I am|I'm) ([^.!?]*)[.!?\s]*$`,
			"Did you come to me because you are $1?",
			"How long have you been $1?",
			"How do you feel about being $1?`",
			"How does being $1 make you feel?`",
			"Do you enjoy being $1?`",
			"Why do you say that you're $1?`",
			"Why do you think you're $1?"},

		{`(?i)^\s*Are you\b([^.!?]*)[.!?\s]*$`,
			"Why does it matter whether I am $1?",
			"Would you prefer it if I were not $1?",
			"Perhaps you believe I am $1.",
			"I may be $1 -- what do you think?"},

		{`(?i)^\s*What ([^.!?]*)[.!?\s]*$`,
			"Why do you ask?",
			"How would an answer to that help you?",
			"What do you think?"},

		{`(?i)^\s*How .*$`,
			"How do you suppose?",
			"Perhaps you can answer your own question.",
			"What is it you're really asking?"},

		{`(?i)^\s*Because ([^.!?]*)[.!?\s]*$`,
			"Is that the real reason?",
			"What other reasons come to mind?",
			"Does that reason apply to anything else?",
			"If $1, what else must be true?"},

		{`(?i)^.*\bsorry\b.*$`,
			"There are many times when no apology is needed.",
			"What feelings do you have when you apologize?"},

		{`(?i)^\s*Hello.*$`,
			"Hello, I'm glad we have a chance to chat.",
			"Hi there, how are you today?",
			"Hello, how are you feeling today?",
			"Alright, mate?"},

		{`(?i)^\s*I think ([^.!?]*)[.!?\s]*$`,
			"Do you doubt $1?",
			"Do you really think so?",
			"But you're not sure $1?"},

		{`(?i)^.*friend.*$`,
			"Tell me more about your friends.",
			"When you think of a friend, what comes to mind?",
			"Did you have many childhood friends?"},
	}

}
