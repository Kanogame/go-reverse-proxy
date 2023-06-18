package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveFolder)
	fmt.Println("Http Server Started and listening at: 12312")
	http.ListenAndServe(":12312", nil)
}

func serveFolder(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/")
}
