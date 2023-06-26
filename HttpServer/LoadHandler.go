package httpserver

import (
	utils "main/Utils"
	"net/http"
	"net/http/httputil"
	"sync/atomic"
)

func (Location *LoadLocations) HandleLoad() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peer := Location.GetNextPeer()
		if peer != nil {
			peer.ReverseProxy.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
	}
}

func (Location *LoadLocations) StartProxyServers() {
	for i, location := range Location.EndPoints {
		proxy := httputil.NewSingleHostReverseProxy(location.URL)
		Location.EndPoints[i] = utils.LoadServer{URL: Location.EndPoints[i].URL, Alive: true, ReverseProxy: proxy}
	}
}

func (Location *LoadLocations) NextIndex() int {
	return int(atomic.AddUint64(&Location.current, uint64(1)) % uint64(len(Location.EndPoints)))
}

func (Location *LoadLocations) GetNextPeer() *utils.LoadServer {
	next := Location.NextIndex()
	l := len(Location.EndPoints) + next
	for i := next; i < l; i++ {
		idx := i % len(Location.EndPoints)
		if Location.EndPoints[idx].IsAlive() {
			if i != next {
				atomic.StoreUint64(&Location.current, uint64(idx))
			}
			return &Location.EndPoints[idx]
		}
	}
	return nil
}

func (Server *LoadServer) SetAlive(alive bool) {
	Server.Alive = alive
}

func (Server *LoadServer) IsAlive() (alive bool) {
	alive = Server.Alive
	return
}
