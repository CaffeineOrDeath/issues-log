package main

import (
	"fmt"      // debugging
	"net/http" // basic server handling

	"github.com/a-h/templ" // templ
)

func main() {

	port := ":8080"

	test := templ.Component(Test())

	http.Handle("/", templ.Handler(test))
	// implemented, no dice
	http.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./styles"))))
	fmt.Println("Starting connection...")
	http.ListenAndServe(port, nil)
}
