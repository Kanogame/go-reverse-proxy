package httpserver

import (
	"fmt"
	utils "main/Utils"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func StartHttpServer(port string, locations *utils.Locations, config *utils.Http) {
	locationHandler(locations, config)
	fmt.Println("Http server started and listening at: ", port)
	err := http.ListenAndServe(":"+port, nil)
	utils.HandleAppError(err)
}

func locationHandler(locations *utils.Locations, config *utils.Http) {
	for _, staticServer := range *locations.Static {
		r := &StaticLocations{&staticServer}
		http.HandleFunc(staticServer.WebPath, r.HandleStatic(config))
	}
	fmt.Println("All static servers started")
	for _, proxyServer := range *locations.Proxy {
		r := &ProxyLocations{&proxyServer}
		go http.HandleFunc(proxyServer.WebPath, r.HandleProxy(config))
	}
	fmt.Println("All reverse proxy servers started")
	for _, loadServer := range *locations.Load {
		r := &LoadLocations{&loadServer}
		r.StartProxyServers()
		go http.HandleFunc(loadServer.WebPath, r.HandleLoad(config))
	}
	fmt.Println("All load servers started")
}
