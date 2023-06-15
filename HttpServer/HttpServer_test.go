package httpserver

import (
	utils "main/Utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleStatic(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:80/", nil)
	w := httptest.NewRecorder()
	test := &[]StaticLocations{
		&utils.StaticLocations{
			WebPath:  "/",
			FilePath: "./static/"},
	}

	http.HandleFunc(test.WebPath, test.HandleStatic())

	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
}
