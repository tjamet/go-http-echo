package main

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	const expectedBody = `{
		"URL": {
		  "Scheme": "http",
		  "Opaque": "",
		  "User": null,
		  "Host": "localhost:8080",
		  "Path": "/some/path",
		  "RawPath": "",
		  "ForceQuery": false,
		  "RawQuery": "",
		  "Fragment": ""
		},
		"Method": "GET",
		"Proto": "HTTP/1.1",
		"Header": {
		  "Someheader": [
			"value"
		  ]
		},
		"Host": "localhost:8080",
		"Body": "body"
	  }`
	r := httptest.NewRequest("GET", "http://localhost:8080/some/path", strings.NewReader("body"))
	r.Header.Set("SomeHeader", "value")
	w := httptest.NewRecorder()
	handler(w, r)
	assert.JSONEq(t, expectedBody, w.Body.String())
}
