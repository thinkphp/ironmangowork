package main

import (
  "fmt"
  "net/http"  
  "os"
)

func getPort() string {

	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {

    http.HandleFunc("/", handler)
    http.ListenAndServe(getPort(), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

     fmt.Fprintf(w, "Hello World!")
}
