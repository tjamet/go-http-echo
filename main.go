package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type R struct {
	URL    *url.URL
	Method string
	Proto  string
	Header http.Header
	Host   string
	body   string
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header()[k] = v
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("X-Request-Read-Error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(&R{r.URL, r.Method, r.Proto, r.Header, r.Host, string(b)})
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}
