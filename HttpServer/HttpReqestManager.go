package httpserver

import (
	"fmt"
	utils "main/Utils"
	"net/http"
)

type StaticLocations struct {
	*utils.StaticLocations
}

func StartHttpServer(port string, locations *utils.Locations) {
	locationHandler(locations)
	http.ListenAndServe(":"+port, nil)
	fmt.Println("Http Server Started and listening at: ", port)
}

func locationHandler(locations *utils.Locations) {
	for _, staticServer := range *locations.Static {
		r := &StaticLocations{&staticServer}
		http.HandleFunc(staticServer.WebPath, r.HandleStatic())
	}
}
