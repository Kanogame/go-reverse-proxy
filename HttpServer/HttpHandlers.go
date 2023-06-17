package httpserver

import (
	utils "main/Utils"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type StaticLocations struct {
	*utils.StaticLocations
}

type ProxyLocations struct {
	*utils.ProxyLocations
}

func (Location *StaticLocations) HandleStatic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, Location.FilePath)
	}
}

func (Location *ProxyLocations) HandleProxy() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, _ := url.Parse(Location.EndPoint)

		proxy := httputil.NewSingleHostReverseProxy(url)

		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host

		proxy.ServeHTTP(w, r)
	}
}
