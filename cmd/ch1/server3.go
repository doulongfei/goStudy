package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

}

func main() {
	http.HandleFunc("/", handler3)
	http.HandleFunc("/gif", func(writer http.ResponseWriter, request *http.Request) {
		lissajousMe(writer)
	})
	http.HandleFunc("/svg", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "imgage/svg+xml")

	})
	log.Fatal(http.ListenAndServe("localhost:9001", nil))
}
