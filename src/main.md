package main

func main() {}

/*
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	fmt.Println("Started:", startTime)

	defer func() {
		elapsed := time.Since(startTime)

		fmt.Println("Ended:", time.Now())
		fmt.Println("Elapsed Seconds:", elapsed.Seconds())
	}()

	f, err := os.Open("/var/log/log-reader-go/access.log")

	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {

		buf := make([]byte, 1024)

		n, err := r.Read(buf)

		buf = buf[:n]

		if n == 0 {
			if err != nil {
				if err == io.EOF {
					break
				}

				panic(err.Error())
			}
			// return err
		}

		linesPool := sync.Pool{
			New: func() interface{} {
				lines := make([]byte, 500*1024)
				return lines
			}}

		stringPool := sync.Pool{
			New: func() interface{} {
				lines := ""
				return lines
			}}

		slicePool := sync.Pool{
			New: func() interface{} {
				lines := make([]string, 100)
				return lines
			}}

		r := bufio.NewReader(f)

		var wg sync.WaitGroup

		for {
			buf := linesPool.Get().([]byte)
			n, err := r.Read(buf)
			buf = buf[:n]

			if n == 0 {
				if err != nil {
					if err == io.EOF {
						break
					}

					panic(err.Error())
				}
			}

			nextUntilNewLine, err := r.ReadBytes('\n')

			if err != io.EOF {
				buf = append(buf, nextUntilNewLine...)
			}

			wg.Add(1)

			go func() {
				ProcessChunk(buf, &linesPool, &stringPool, &slicePool)
				// fmt.Println(string(buf), stringPool, slicePool)
				wg.Done()
			}()

			wg.Wait()

		}

	}

}

func ProcessChunk(chunk []byte, linesPool *sync.Pool, stringPool *sync.Pool, slicePool *sync.Pool) {

}
*/
