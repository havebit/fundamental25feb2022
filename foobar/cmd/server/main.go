package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"gitlab.com/cjexpress/tildi/signac/learngo/foobar"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/foobar/", foobarHandler)

	http.Handle("/foobar2/", http.HandlerFunc(foobarHandler))

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func foobarHandler(rw http.ResponseWriter, r *http.Request) {
	param := r.RequestURI[8:]
	result := foobar.SayAny(param)

	m := map[string]string{
		"foobar": result,
	}

	if err := json.NewEncoder(rw).Encode(&m); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		io.WriteString(rw, err.Error())
		return
	}
}
