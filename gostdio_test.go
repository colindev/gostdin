package gostdio

import (
	"fmt"
	"io"
)

func Example() {

	r, w := io.Pipe()
	go func() {
		w.Write([]byte("a\n"))
		w.Write([]byte("b\r\n"))
		w.Write([]byte("c\n"))
		w.Close()
	}()

	err := Run(r, func(p []byte) error {
		fmt.Println(string(p))
		return nil
	})

	fmt.Println(err)
	// output:
	// a
	// b
	// c
	// EOF
}
