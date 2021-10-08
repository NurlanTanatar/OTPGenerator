package shortner

import (
	"encoding/json"
	"io"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}

		fallback.ServeHTTP(w, r)
	}
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc that will attempt to map any paths to their
// corresponding URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
// [
//    {
//      "path": "/some-path",
//      "url": "https://www.some-url.com/demo"
//    }
// ]
func JSONHandler(r io.Reader, fallback http.Handler) (http.HandlerFunc, error) {
	decoder := json.NewDecoder(r)
	pathURLs, err := decode(decoder)
	if err != nil {
		return nil, err
	}
	pathToUrls := buildMap(pathURLs)

	mapHandler := MapHandler(pathToUrls, fallback)
	return mapHandler, nil
}

type pathURL struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

type decoder interface {
	Decode(v interface{}) error
}

func decode(d decoder) ([]pathURL, error) {
	var pu []pathURL
	for {
		err := d.Decode(&pu)
		if err == io.EOF {
			return pu, nil
		} else if err != nil {
			return nil, err
		}
	}
}

func buildMap(pathURLs []pathURL) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range pathURLs {
		pathToUrls[pu.Path] = pu.URL
	}

	return pathToUrls
}
