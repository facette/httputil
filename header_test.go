package httputil

import (
	"net/http"
	"testing"
)

func Test_GetContentType(t *testing.T) {
	if result := GetContentType(nil); result != "" {
		t.Logf("\nExpected %q\nbut got  %q", "", result)
		t.Fail()
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	if result := GetContentType(req); result != "" {
		t.Logf("\nExpected %q\nbut got  %q", "", result)
		t.Fail()
	}

	req.Header.Add("Content-Type", "application/json")
	if result := GetContentType(req); result != "application/json" {
		t.Logf("\nExpected %q\nbut got  %q", "application/json", result)
		t.Fail()
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if result := GetContentType(req); result != "application/json" {
		t.Logf("\nExpected %q\nbut got  %q", "application/json", result)
		t.Fail()
	}

	resp := &http.Response{Header: http.Header{}}

	if result := GetContentType(resp); result != "" {
		t.Logf("\nExpected %q\nbut got  %q", "", result)
		t.Fail()
	}

	resp.Header.Add("Content-Type", "application/json")
	if result := GetContentType(resp); result != "application/json" {
		t.Logf("\nExpected %q\nbut got  %q", "application/json", result)
		t.Fail()
	}

	resp.Header.Set("Content-Type", "application/json; charset=utf-8")
	if result := GetContentType(resp); result != "application/json" {
		t.Logf("\nExpected %q\nbut got  %q", "application/json", result)
		t.Fail()
	}
}
