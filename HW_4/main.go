package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/NurlanTanatar/golang_hw/tree/main/HW_4/shortner"
)

const (

	// JSONFlag is used to set a file for the questions
	JSONFlag = "json"
	// JSONFlagValue is the value used when no JSONFlag is provided
	JSONFlagValue = "map.json"
	// JSONFlagUsage is the help string for the JSONFlag
	JSONFlagUsage = "URLs file in json format"
)

// Flagger is an interface for configuring various flags
type Flagger interface {
	StringVar(p *string, name, value, usage string)
}

type shortnerFlagger struct{}

func (uf *shortnerFlagger) StringVar(p *string, name, value, usage string) {
	flag.StringVar(p, name, value, usage)
}

var json string

// ConfigFlags will configure the flags used by the application
func ConfigFlags(flagger Flagger) {
	flagger.StringVar(&json, JSONFlag, JSONFlagValue, JSONFlagUsage)
}

func main() {
	flagger := &shortnerFlagger{}
	ConfigFlags(flagger)

	mapHandler := createMapHandler()
	jsonHandler := createJSONHandler(mapHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

// Build the MapHandler using the mux as the fallback
var pathsToUrls = map[string]string{
	"/shortner-godoc": "https://godoc.org/github.com/gophercises/shortner",
	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
}

func createMapHandler() http.HandlerFunc {
	mux := defaultMux()
	return shortner.MapHandler(pathsToUrls, mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func createJSONHandler(fallback http.HandlerFunc) http.HandlerFunc {
	jsonFile, err := os.Open(json)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := shortner.JSONHandler(jsonFile, fallback)
	if err != nil {
		panic(err)
	}

	return jsonHandler
}
