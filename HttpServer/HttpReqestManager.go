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
	fmt.Println("Http Server Started and listening at: ", port)
	http.ListenAndServe(":"+port, nil)
}

func locationHandler(locations *utils.Locations) {
	for _, staticServer := range *locations.Static {
		r := &StaticLocations{&staticServer}
		go http.HandleFunc(staticServer.WebPath, r.HandleStatic())
	}
}
