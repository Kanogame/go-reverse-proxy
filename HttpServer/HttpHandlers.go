package httpserver

import (
	"net/http"
)

func (Location *StaticLocations) HandleStatic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, Location.FilePath)
	}
}
