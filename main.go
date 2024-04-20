package main

import (
	"fmt"
	"net/http"
)

func getHello() string {
	return "hello world"
}

func mountainHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Anonymous"
	}

	// レスポンスを返す
	fmt.Fprintf(w, "Hello, %s! The server is healthy.", username)
}

func main() {
	fmt.Println(getHello())
	http.HandleFunc("/mountain", mountainHandler)

	fmt.Println("server is runnning")
	http.ListenAndServe(":8080", nil)
}
