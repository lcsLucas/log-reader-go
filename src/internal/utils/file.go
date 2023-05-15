package utils

import (
	"fmt"
	"io"
	"os"
)

func ReadLine(f *os.File, offset uint64, back bool) (b []byte, err error) {
	lastLineSize := 0
	var s string

	for {
		b := make([]byte, 1)
		n, err := f.ReadAt(b, int64(offset))

		if err != nil && !back && err != io.EOF {
			return []byte{}, err
		} else if back && err == io.EOF {
			fmt.Print("back", "EOF")
			continue
		}

		char := string(b[0])

		s += char

		if char == "\n" && lastLineSize != 0 {
			break
		}

		if back {
			offset--
		} else {
			offset++
		}

		lastLineSize += n

	}

	rns := []byte(s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return rns, err
}
