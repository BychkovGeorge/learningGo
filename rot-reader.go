package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case b >= 'a' && b <= 'z':
		return 'a' + (b-'a'+13)%26
	case b >= 'A' && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	}
	return b
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	a, err := r.r.Read(p)
	for i, b := range p[:a] {
		p[i] = rot13(b)
	}
	return a, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
