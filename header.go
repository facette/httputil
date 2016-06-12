package httputil

import (
	"net/http"
	"strings"
)

// GetContentType returns the HTTP 'Content-Type' header value.
func GetContentType(v interface{}) string {
	var header http.Header

	switch v.(type) {
	case *http.Request:
		header = v.(*http.Request).Header
	case *http.Response:
		header = v.(*http.Response).Header
	default:
		return ""
	}

	ct := header.Get("Content-Type")

	index := strings.Index(ct, ";")
	if index != -1 {
		return ct[:index]
	}

	return ct
}
