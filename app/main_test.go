package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/a-h/templ"
	"github.com/crispyfishtech/crispyfish-demo/views"
)

func TestGetIndex(t *testing.T) {
	ts := httptest.NewServer(templ.Handler(views.Index("Crispyfish Demo")))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)

	}

	if res.StatusCode != 200 {
		t.Fatal("expected 200 status")
	}
}

func TestGetPing(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(ping))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)

	}

	if res.StatusCode != 200 {
		t.Fatal("expected 200 status")
	}
}
