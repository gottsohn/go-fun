// main_test.go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	a := App{}
	a.Initializer("localhost", "movies_test")
	clearDatabase()
	// a.Runr(3000)
	code := m.Run()
	os.Exit(code)
}

func clearDatabase() {
	a.dao.RemoveAll()
}

func TestSanityCheck(t *testing.T) {
	t.Logf("Sanity Check")
	var testString = "expect 1 to be 1"
	if 1 == 1 {
		t.Logf(testString)
	} else {
		t.Errorf(testString)
	}
}

func ToTestGetNonExistentMovie(t *testing.T) {
	req, _ := http.NewRequest("GET", "/movies/111111111111111111111111", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Movie not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
