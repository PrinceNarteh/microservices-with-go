package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Goodby World!")
	})

	http.ListenAndServe(":4000", nil)
}
