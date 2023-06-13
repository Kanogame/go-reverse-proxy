package httpserver

import (
	"fmt"
	utils "main/Utils"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

func StartHttpServer(port int, locations *utils.Locations) {
	var c = cors.New(cors.Options{
		AllowedOrigins: []string{ /*all*/ },
	})

	responseHandler := func(w http.ResponseWriter, r *http.Request) {
		return HttpHandler(w, r, locations)
	}

	handler := http.HandlerFunc(responseHandler())
	fmt.Println("Http Server Started and listening at: ", port)
	http.ListenAndServe(":"+strconv.Itoa(port), c.Handler(handler))
}

func (locations *utils.Locations) HttpHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if r.Method == "GET" {
		//TODO
		//http.ServeFile(w, r, path)
	} else if r.Method == "POST" {
		//TODO
	}
}
