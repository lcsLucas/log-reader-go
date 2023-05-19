package file

import (
	"bufio"
	"context"
	"io"
	"log-reader-go/internal/utils/regex"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
)

var m sync.Mutex

var sem *semaphore.Weighted

func ProcessFile(ctx context.Context, f *os.File, startTime *time.Time, endTime *time.Time) error {

	linesPool := sync.Pool{
		New: func() interface{} {
			l := new([]byte)
			*l = make([]byte, 15*1024*1024) //20MB

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

	sem = semaphore.NewWeighted(150)

	for {

		select {
		case <-ctx.Done():
			return nil
		default:
		}

		buf := linesPool.Get().(*[]byte)

		b := *buf
		n, err := r.Read(b)
		*buf = b[:n]

		if err != nil && err != io.EOF {
			log.Warning(err, n)
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

		sem.Acquire(ctx, 1)

		go func() {

			defer func() {
				wg.Done()
				sem.Release(1)
			}()

			ParserChunk(buf, &linesPool, &stringPool, startTime, endTime)
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

	logsSlice := splitLines(*logs)
	// strings.Split(*logs, "\n")

	stringPool.Put(logs)

	n := len(logsSlice)

	for i := 0; i < n; i++ {
		wg2.Add(1)
		go func(l string) {
			defer wg2.Done()

			if len(l) != 0 {
				_, err := regex.LogParse(l)

				if err != nil {
					m.Lock()

					log.Error(err.Error())
					log.Info(l)

					m.Unlock()
				}
			}
		}(logsSlice[i])
	}

	wg2.Wait()
	logsSlice = nil
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}
