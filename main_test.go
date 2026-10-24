package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	assert.JSONEq(t, expectedBody, withoutNewURLFields(t, w.Body.String()))
}

// withoutNewURLFields drops the url.URL fields that newer Go versions
// serialise (OmitHost since 1.19, RawFragment since 1.15), so the expectation
// holds on every Go version.
func withoutNewURLFields(t *testing.T, body string) string {
	t.Helper()
	var decoded map[string]interface{}
	require.NoError(t, json.Unmarshal([]byte(body), &decoded))
	url, ok := decoded["URL"].(map[string]interface{})
	require.True(t, ok, "URL is not an object in %s", body)
	delete(url, "OmitHost")
	delete(url, "RawFragment")
	out, err := json.Marshal(decoded)
	require.NoError(t, err)
	return string(out)
}
