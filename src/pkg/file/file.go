package file

import (
	"bufio"
	"context"
	"io"
	"log-reader-go/internal/log"
	"log-reader-go/internal/utils/regex"
	"os"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var m sync.Mutex

var sem *semaphore.Weighted

var sem2 chan struct{}

func ProcessFile(ctx context.Context, f *os.File, startTime *time.Time, endTime *time.Time) error {

	linesPool := sync.Pool{
		New: func() interface{} {
			l := new([]byte)
			*l = make([]byte, 500*1024) //20MB

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

	sem = semaphore.NewWeighted(int64(runtime.NumCPU()))

	sem2 = make(chan struct{}, int64(runtime.NumCPU()))

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
			log.Logger.Warning(err, n)
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

			ParserChunk(&sem2, buf, &linesPool, &stringPool, startTime, endTime)
		}()

	}

	wg.Wait()
	return nil
}

func ParserChunk(sem *chan struct{}, chunk *[]byte, linesPool *sync.Pool, stringPool *sync.Pool, startTime *time.Time, endTime *time.Time) {

	var wg2 sync.WaitGroup

	logs := stringPool.Get().(*string)
	*logs = string(*chunk)

	linesPool.Put(chunk)

	logsSlice := splitLines(*logs)

	stringPool.Put(logs)

	n := len(logsSlice)

	for i := 0; i < n; i++ {
		wg2.Add(1)
		*sem <- struct{}{}

		go func(l string) {

			defer func() {
				wg2.Done()
				<-*sem
			}()

			if len(l) != 0 {
				_, err := regex.LogParse(l)

				if err != nil {
					m.Lock()

					log.Logger.Error(err.Error())
					log.Logger.Info(l)

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
