package httpserver

import (
	utils "main/Utils"
	"net/http"
)

func (Location *LoadLocations) HandleLoad() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Location.StartProxyServers()
	}
}

func (Location *LoadLocations) StartProxyServers() {
	for i, location := range Location.EndPoints {
		Location.EndPoints[i] = utils.LoadServer{}
	}
}
