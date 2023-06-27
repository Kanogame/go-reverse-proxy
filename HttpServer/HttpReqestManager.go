package httpserver

import (
	"fmt"
	utils "main/Utils"
	"net/http"
)

func StartHttpServer(port string, locations *utils.Locations) {
	locationHandler(locations)
	fmt.Println("Http Server Started and listening at: ", port)
	err := http.ListenAndServe(":"+port, nil)
	utils.HandleAppError(err)
}

func locationHandler(locations *utils.Locations) {
	for _, staticServer := range *locations.Static {
		r := &StaticLocations{&staticServer}
		go http.HandleFunc(staticServer.WebPath, r.HandleStatic())
	}
	for _, proxyServer := range *locations.Proxy {
		r := &ProxyLocations{&proxyServer}
		go http.HandleFunc(proxyServer.WebPath, r.HandleProxy())
	}
	for _, loadServer := range *locations.Load {
		r := &LoadLocations{&loadServer}
		r.StartProxyServers()
		go http.HandleFunc(loadServer.WebPath, r.HandleLoad())
	}
}
