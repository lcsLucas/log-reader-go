package file

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

func OpenFile(filename string) (f *os.File, err error) {
	f, err = os.Open(filename)
	return
}

func StatFile(f *os.File) (stat fs.FileInfo, err error) {
	stat, err = f.Stat()
	return
}

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

	if back {
		for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
			rns[i], rns[j] = rns[j], rns[i]
		}
	}

	return rns, err
}
