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

     fmt.Fprintf(w, "<h1 style='font-size: 80px;background-color: yellow;color: mediumseagreen'>First App Chelner in Romania with Go! stay tuned koi!</h1><p>Created with Passion by Adrian Statescu</p>")
}
