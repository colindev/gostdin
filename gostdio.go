package gostdio

import (
	"bufio"
	"io"
)

// Run io stream processor from reader
func Run(r io.Reader, fn func([]byte) error) error {

	in := bufio.NewReader(r)

	for {
		line, _, err := in.ReadLine()
		if err != nil {
			return err
		}

		if err := fn(line); err != nil {
			return err
		}
	}
}
