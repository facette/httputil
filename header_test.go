package httputil

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetContentType(t *testing.T) {
	actual, err := GetContentType(nil)
	assert.Equal(t, ErrInvalidInterface, err)
	assert.Empty(t, actual)

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)

	actual, err = GetContentType(req)
	assert.NotNil(t, err)
	assert.Empty(t, actual)

	req.Header.Add("Content-Type", "application/json")
	actual, err = GetContentType(req)
	assert.Nil(t, err)
	assert.Equal(t, "application/json", actual)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	actual, err = GetContentType(req)
	assert.Nil(t, err)
	assert.Equal(t, "application/json", actual)

	resp := &http.Response{Header: http.Header{}}

	actual, err = GetContentType(resp)
	assert.NotNil(t, err)
	assert.Empty(t, actual)

	resp.Header.Add("Content-Type", "application/json")
	actual, err = GetContentType(resp)
	assert.Nil(t, err)
	assert.Equal(t, "application/json", actual)

	resp.Header.Set("Content-Type", "application/json; charset=utf-8")
	actual, err = GetContentType(resp)
	assert.Nil(t, err)
	assert.Equal(t, "application/json", actual)
}
