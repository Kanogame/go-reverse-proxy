package httpserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

func StartHttpServer(port int) {
	var c = cors.New(cors.Options{
		AllowedOrigins: []string{ /*all*/ },
	})

	handler := http.HandlerFunc(HttpHandler)
	fmt.Println("Http Server Started and listening at: ", port)
	http.ListenAndServe(":"+strconv.Itoa(port), c.Handler(handler))
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if r.Method == "GET" {
		//TODO
		//http.ServeFile(w, r, path)
	} else if r.Method == "POST" {
		//TODO
	}
}
