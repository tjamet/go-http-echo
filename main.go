package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
	_, err = w.Write(b)
	if err != nil {
		fmt.Printf("failed to write response: %s", err.Error())
	}
	fmt.Println("path: ", r.RequestURI, "method: ", r.Method)
	fmt.Println("body: ", string(b))
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}
