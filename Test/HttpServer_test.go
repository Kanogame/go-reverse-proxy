package tests

import (
	"io/ioutil"
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

	ts := httptest.NewServer(http.HandlerFunc(test.HandleStatic(nil)))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("expected 200 got %v", res.StatusCode)
	}
}

func TestHandleProxy(t *testing.T) {
	test := &httpserver.ProxyLocations{
		ProxyLocations: &utils.ProxyLocations{
			WebPath:  "/",
			EndPoint: "http://127.0.0.1:12312/"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("<div>hello world</div>")) })
	t.Log("Test Server listening at: 12312")
	go http.ListenAndServe(":12312", nil)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	test.HandleProxy(nil).ServeHTTP(w, req)

	res := w.Result()
	if res.StatusCode != 200 {
		t.Errorf("expected 200 got %v", res.StatusCode)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	utils.HandleAppError(err)
	if string(data) != "<div>hello world</div>" {
		t.Errorf("expected <div>hello world</div> got %v", string(data))
	}
}
