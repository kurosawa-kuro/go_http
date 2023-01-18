package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("Hello world!")
	})

	http.ListenAndServe(":9090", nil)
}
