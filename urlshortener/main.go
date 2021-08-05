package main

import (
	"fmt"
	"log"
	"net/http"

	urlshort "github.com/pashamakhilkumarreddy/gophercises/urlshortener/handlers/urlshort"
)

func main() {
	fmt.Println("Welcome to URLShortener")
	mux := defaultMux()
	pathToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathToUrls, mux)
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/final`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Starting on server http://localhost:5000")
	http.ListenAndServe(":5000", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo!")
}
