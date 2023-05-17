package file

import (
	"bufio"
	"io"
	"log-reader-go/pkg/color"
	"math"
	"os"
	"strings"
	"sync"
	"time"
)

func ProcessFile(f *os.File, startTime *time.Time, endTime *time.Time) error {

	linesPool := sync.Pool{
		New: func() interface{} {
			l := new([]byte)
			*l = make([]byte, 1024*1024)

			return l
		},
	}

	stringPool := sync.Pool{
		New: func() interface{} {
			return new(string)
		},
	}

	r := bufio.NewReader(f)

	var wg sync.WaitGroup

	for {
		buf := linesPool.Get().(*[]byte)

		b := *buf
		n, err := r.Read(b)
		*buf = b[:n]

		if err != nil && err != io.EOF {
			color.PrintYellow(err, " ", n)
			continue
		}

		if n == 0 || err == io.EOF {
			break
		}

		// lê o pedaço até o '\n' e junta com o outro pedaço lido no buf,
		//evitando assim de cortar registros nas leituras por bytes
		readUntillNewline, err := r.ReadBytes('\n')

		if err != io.EOF {
			*buf = append(*buf, readUntillNewline...)
		}

		wg.Add(1)

		go func() {
			ParserChunk(buf, &linesPool, &stringPool, startTime, endTime)
			wg.Done()
		}()

	}

	wg.Wait()
	return nil
}

func ParserChunk(chunk *[]byte, linesPool *sync.Pool, stringPool *sync.Pool, startTime *time.Time, endTime *time.Time) {
	var wg2 sync.WaitGroup

	logs := stringPool.Get().(*string)
	*logs = string(*chunk)

	linesPool.Put(chunk)

	logsSlice := strings.Split(*logs, "\n")

	stringPool.Put(logs)

	chunkSize := 100
	n := len(logsSlice)
	noOfThread := n / chunkSize

	if n%chunkSize != 0 {
		noOfThread++
	}

	for i := 0; i < noOfThread; i++ {
		wg2.Add(1)

		go func(s int, e int) {
			defer wg2.Done()

			for i := s; i < e; i++ {
				text := logsSlice[i]

				if len(text) == 0 {
					continue
				}
			}

		}(i*chunkSize, int(math.Min(float64((i+1)*chunkSize), float64(len(logsSlice)))))
	}

	wg2.Wait()
	logsSlice = nil

}
