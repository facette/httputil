package httputil

import (
	"bytes"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	jsonData = "{\"key1\":1,\"key2\":2}\n"
	jsonMap  = map[string]int{"key1": 1, "key2": 2}
)

func Test_BindJSON(t *testing.T) {
	var actual map[string]int

	req, err := http.NewRequest("GET", "/", strings.NewReader(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer req.Body.Close()
	req.Header.Set("Content-Type", "application/json")

	err = BindJSON(req, &actual)
	assert.Nil(t, err)
	assert.Equal(t, jsonMap, actual)

	resp := &http.Response{Header: http.Header{}, Body: ioutil.NopCloser(strings.NewReader(jsonData))}
	resp.Header.Add("Content-Type", "application/json")

	err = BindJSON(resp, &actual)
	assert.Nil(t, err)
	assert.Equal(t, jsonMap, actual)
}

func Test_BindJSON_Invalid(t *testing.T) {
	var actual map[string]int

	req, err := http.NewRequest("GET", "/", strings.NewReader(`{"invalid":`))
	if err != nil {
		t.Fatal(err)
	}
	defer req.Body.Close()
	req.Header.Set("Content-Type", "application/json")

	err = BindJSON(req, &actual)
	assert.NotNil(t, err)
	assert.Nil(t, actual)

	resp := &http.Response{Header: http.Header{}, Body: ioutil.NopCloser(strings.NewReader(`{"invalid":`))}
	resp.Header.Add("Content-Type", "application/json")

	err = BindJSON(resp, &actual)
	assert.NotNil(t, err)
	assert.Nil(t, actual)

	err = BindJSON(nil, &actual)
	assert.Equal(t, ErrInvalidInterface, err)
	assert.Nil(t, actual)
}

func Test_BindJSON_Invalid_ContentType(t *testing.T) {
	var actual map[string]int

	req, err := http.NewRequest("GET", "/", strings.NewReader(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer req.Body.Close()
	req.Header.Add("Content-Type", "text/plain")

	err = BindJSON(req, &actual)
	assert.Equal(t, ErrInvalidContentType, err)
	assert.Nil(t, actual)

	resp := &http.Response{Header: http.Header{}, Body: ioutil.NopCloser(strings.NewReader(jsonData))}
	resp.Header.Add("Content-Type", "text/plain")

	err = BindJSON(resp, &actual)
	assert.Equal(t, ErrInvalidContentType, err)
	assert.Nil(t, actual)
}

func Test_WriteJSON(t *testing.T) {
	b := bytes.NewBuffer(nil)
	w := testWriter{b}

	err := WriteJSON(w, jsonMap, http.StatusOK)
	assert.Nil(t, err)
	assert.Equal(t, jsonData, b.String())
}

func Test_WriteJSON_Invalid(t *testing.T) {
	b := bytes.NewBuffer(nil)
	w := testWriter{b}

	err := WriteJSON(w, math.NaN, http.StatusOK)
	assert.NotNil(t, err)
}

type testWriter struct {
	io.Writer
}

func (testWriter) Header() http.Header {
	return http.Header{}
}

func (testWriter) WriteHeader(int) {}
