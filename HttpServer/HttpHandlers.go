package httpserver

import (
	"fmt"
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

type LoadLocations struct {
	*utils.LoadLocations
}

func (Location *StaticLocations) HandleStatic(config *utils.Http) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if config != nil && r.URL.Path != Location.WebPath {
			http.ServeFile(w, r, config.File404)
			return
		}
		http.ServeFile(w, r, Location.FilePath)
	}
}

func (Location *ProxyLocations) HandleProxy(config *utils.Http) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if config != nil && r.URL.Path != Location.WebPath {
			http.ServeFile(w, r, config.File404)
			return
		}
		url, _ := url.Parse(Location.EndPoint)

		proxy := httputil.NewSingleHostReverseProxy(url)

		fmt.Println(r.URL.Host)
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

		proxy.ServeHTTP(w, r)
	}
}
