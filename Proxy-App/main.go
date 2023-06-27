package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", serveFolder)
	fmt.Println("Http Server Started and listening at: 12312")
	http.ListenAndServe(":"+os.Args[2], nil)
}

func serveFolder(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, os.Args[1])
}
