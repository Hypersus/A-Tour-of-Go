// https://go.dev/tour/methods/22

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (myreader MyReader) Read(b []byte) (int, error) {
	for i,_ := range b {
		b[i] = 'A'
	}
	return len(b),nil
}

func main() {
	reader.Validate(MyReader{})
}
