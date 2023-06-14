package httpserver

import (
	"fmt"
	utils "main/Utils"
	"net/http"
)

type StaticLocations struct {
	*utils.StaticLocations
}

func StartHttpServer(port string, locations utils.Locations) {
	locationHandler(locations)
	http.ListenAndServe(":"+port, nil)
	fmt.Println("Http Server Started and listening at: ", port)
}

func locationHandler(locations utils.Locations) {
	var static []StaticLocations = &locations.static
	for _, staticServer := range static {
		http.HandleFunc(staticServer.WebPath, staticServer.HandleStatic)
	}
}
