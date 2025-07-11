package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var responses = []string{
	"It is certain.",
	"Ask again later.",
	"Don't count on it.",
	"Yes – definitely.",
	"Outlook not so good.",
	"Most likely.",
	"My sources say no.",
	"Cannot predict now.",
	"You may rely on it.",
	"Very doubtful.",
}

func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	response := responses[rand.Intn(len(responses))]

	fmt.Fprintf(w, "<h1>🎱 Magic 8-Ball 🎱</h1><p><em>%s</em></p><p>Refresh for a new answer.</p>", response)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
