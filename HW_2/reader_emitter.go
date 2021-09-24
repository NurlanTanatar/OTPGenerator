package main

// install go tour libraries and files first go install golang.org/x/website/tour@latest
import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return 1, nil
}
func checker() {
	reader.Validate(MyReader{})
}
