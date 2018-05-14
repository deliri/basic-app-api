package main

import (
	"fmt"
	"github.com/maxdobeck/gatekeeper/authentication"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Consider adding to golang examples for httptest
func TestLogin(t *testing.T) {
	bodyReader := strings.NewReader(`{"email": "WrongEmail@email.com", "password": "wrongPassword"}`)
	req, err := http.NewRequest("POST", "/login", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	gatekeeper.login(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
