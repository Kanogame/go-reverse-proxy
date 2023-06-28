package httpserver

import (
	"fmt"
	utils "main/Utils"
	"net/http"
)

func StartHttpServer(port string, locations *utils.Locations) {
	locationHandler(locations)
	fmt.Println("Http server started and listening at: ", port)
	err := http.ListenAndServe(":"+port, nil)
	utils.HandleAppError(err)
}

func locationHandler(locations *utils.Locations) {
	for _, staticServer := range *locations.Static {
		r := &StaticLocations{&staticServer}
		go http.HandleFunc(staticServer.WebPath, r.HandleStatic())
	}
	fmt.Println("All static servers started")
	for _, proxyServer := range *locations.Proxy {
		r := &ProxyLocations{&proxyServer}
		go http.HandleFunc(proxyServer.WebPath, r.HandleProxy())
	}
	fmt.Println("All reverse proxy servers started")
	for _, loadServer := range *locations.Load {
		r := &LoadLocations{&loadServer}
		r.StartProxyServers()
		go http.HandleFunc(loadServer.WebPath, r.HandleLoad())
	}
	fmt.Println("All load servers started")
}
