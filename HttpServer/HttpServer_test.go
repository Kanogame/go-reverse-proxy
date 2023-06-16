package httpserver

import (
	utils "main/Utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleStatic(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index.html", nil)
	w := httptest.NewRecorder()
	test := &StaticLocations{
		&utils.StaticLocations{
			WebPath:  "/",
			FilePath: "./static/"},
	}

	handler := http.HandlerFunc(test.HandleStatic())
	handler(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Error(w.Result())
		t.Errorf("expected a %d, instead got: %d", http.StatusOK, w.Result().StatusCode)
	}
}
