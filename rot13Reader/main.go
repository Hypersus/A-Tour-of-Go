// https://go.dev/tour/methods/23

package main

import (
	"io"
	"os"
	"strings"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (rR rot13Reader) Read(b []byte) (n int, err error) {
	for {
		n, err = rR.r.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Println("Before")
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b[:n])
		fmt.Printf("b[:n] = %q\n", b[:n])
		for i:=0;i<len(b);i++ {
			isCap := b[i] >= 'A' && b[i]<='Z'
			isSmall := b[i] >= 'a' && b[i] <= 'z'
			if isCap {
				b[i]=(b[i]-'A'+13)%26+'A'
			} else if isSmall {
				b[i]=(b[i]-'a'+13)%26+'a'
			}
		}
		fmt.Println("After")
        fmt.Printf("n = %v err = %v b = %v\n", n, err, b[:n])
		fmt.Printf("b[:n] = %q\n", b[:n])
	}
	return n,err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, r)
}
