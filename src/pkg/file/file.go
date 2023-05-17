package file

import (
	"bufio"
	"io"
	"log-reader-go/pkg/color"
	"os"
	"sync"
	"time"
)

func ProcessFile(f *os.File, startTime *time.Time, endTime *time.Time) error {

	linesPool := sync.Pool{
		New: func() interface{} {
			lines := make([]byte, 250*1024)
			return lines
		},
	}

	stringPool := sync.Pool{
		New: func() interface{} {
			lines := ""
			return lines
		},
	}

	r := bufio.NewReader(f)

	var wg sync.WaitGroup

	for {
		buf := linesPool.Get().([]byte)

		n, err := r.Read(buf)
		buf = buf[:n]

		if err != nil && err != io.EOF {
			color.PrintYellow(err, " ", n)
			continue
		}

		if n == 0 || err == io.EOF {
			break
		}

		// lê um pedaço até encontrar '\n' para juntar no buf, para não acontecer de ter lido apenas uma parte da linha, cortando o registro.
		readUntillNewline, err := r.ReadBytes('\n')

		// junta as partes lida com buffer e read bytes
		if err != io.EOF {
			buf = append(buf, readUntillNewline...)
		}

		wg.Add(1)
		defer wg.Wait()

		go func() {
			ParserChunk(buf, &linesPool, &stringPool, startTime, endTime)
			wg.Done()
		}()

	}

	return nil
}

func ParserChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, startTime *time.Time, endTime *time.Time) {
	// color.PrintPurple(string(chunk[:1]), " | ")
}
