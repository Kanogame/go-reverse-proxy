package utils

import (
	"net/http/httputil"
	"net/url"
)

type Http struct {
	Port      string
	LogFolder string
}

type UndefinedLocation struct {
	WebPath string
	Utype   string
	Path    string
}

type StaticLocations struct {
	WebPath  string
	FilePath string
}

type ProxyLocations struct {
	WebPath  string
	EndPoint string
}

type LoadLocations struct {
	WebPath   string
	Current   uint64
	EndPoints []LoadServer
}

type LoadServer struct {
	URL          *url.URL
	Alive        bool
	ReverseProxy *httputil.ReverseProxy
}

func (Server *LoadServer) SetAlive(alive bool) {
	Server.Alive = alive
}

func (Server *LoadServer) IsAlive() (alive bool) {
	alive = Server.Alive
	return
}

type Locations struct {
	Static *[]StaticLocations
	Proxy  *[]ProxyLocations
	Load   *[]LoadLocations
}
