package tests

import (
	"log"
	httpserver "main/HttpServer"
	utils "main/Utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleStatic(t *testing.T) {
	test := &httpserver.StaticLocations{
		StaticLocations: &utils.StaticLocations{
			WebPath:  "/",
			FilePath: "../static/"},
	}

	ts := httptest.NewServer(http.HandlerFunc(test.HandleStatic()))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("expected 200 got %v", res.StatusCode)
	}
}
