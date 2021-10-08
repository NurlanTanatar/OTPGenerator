package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/NurlanTanatar/golang_hw/tree/main/HW_4/shortner"
)

const DEFAULTFILE = "map.json"

func main() {

	mux := defaultMux()
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := shortner.MapHandler(pathsToUrls, mux)

	// pairProducer will be used in the MainHandler
	var pairProducer shortner.PairProducer

	// mainHandler will be used as in ListenAndServe
	mainHandler, err := shortner.MainHandler(pairProducer, mapHandler)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mainHandler)
}

// getContent opens file and returns urlshort.Content
func getContent(file string) (shortner.Content, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return shortner.Content(content), nil
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
