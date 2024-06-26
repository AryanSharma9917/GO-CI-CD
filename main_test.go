package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
)

func TestPrintName(t *testing.T) {
	actualname := "Ram"
	name := PrintName(actualname)
	if name != actualname {
		t.Errorf("Expected: %s, Actual: %s", actualname, name)
	}

	actualname = "Aryan"
	name = PrintName(actualname)
	if name != actualname {
		t.Errorf("Expected: %s, Actual: %s", actualname, name)
	}

	actualname = "Charles"
	name = PrintName(actualname)
	if name != actualname {
		t.Errorf("Expected: %s, Actual: %s", actualname, name)
	}

	actualname = "SHAYM"
	name = PrintName(actualname)
	if name != actualname {
		t.Errorf("Expected: %s, Actual: %s", actualname, name)
	}

	actualname = "anirban"
	name = PrintName(actualname)
	if name != actualname {
		t.Errorf("Expected: %s, Actual: %s", actualname, name)
	}

}

func TestApi(t *testing.T) {
	m := martini.Classic()

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + PrintName(params["name"]) + "\n"
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello/Aryan", nil)

	m.ServeHTTP(res, req)

	if res.Body.String() != "Hello Aryan\n" {
		t.Errorf("Got: %s", res.Body.String())
	}

	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/hello/Charles", nil)

	m.ServeHTTP(res, req)

	if res.Body.String() != "Hello Charles\n" {
		t.Errorf("Got: %s", res.Body.String())
	}

	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/hello/James", nil)

	m.ServeHTTP(res, req)

	if res.Body.String() != "Hello James\n" {
		t.Errorf("Got: %s", res.Body.String())
	}
}
