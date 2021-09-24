package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i, chr := range p {
		switch {
		case chr >= 65 && chr < 90:
			chr = 65 + ((chr - 52) % 26)
		case chr >= 97 && chr <= 122:
			chr = 97 + ((chr - 84) % 26)
		}
		p[i] = chr
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
