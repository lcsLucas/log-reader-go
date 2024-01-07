package main

import (
	"context"
	"log-reader-go/internal/config"
	"log-reader-go/internal/log"
	"log-reader-go/internal/utils/args"
	"log-reader-go/internal/utils/env"
	"log-reader-go/internal/utils/file"
	"log-reader-go/internal/validate"
	file2 "log-reader-go/pkg/file"
	"os"
	"os/signal"
	"time"
)

var (
	startTimeExec time.Time
	cLog          config.LogFile
)

func catchError(err error) {
	if err != nil {
		log.Logger.Error(err.Error())
		os.Exit(-1)
	}
}

func checkDateLog(f *os.File, t *time.Time, offset int64) error {
	if t != nil {
		lineRead, err := file.ReadLine(f, uint64(offset), true)

		if err != nil {
			return err
		}

		if offset == 0 {
			err = validate.ValidateDateRangeLogs(cLog.LogStartTime, lineRead, false)
		} else {
			err = validate.ValidateDateRangeLogs(cLog.LogStartTime, lineRead, true)
		}

		return err
	}

	return nil
}

func main() {
	var err error

	defer end()

	// env load
	err = env.Load()
	catchError(err)

	// elastic.Eae()

	// flags read
	err = args.Read(&cLog)
	catchError(err)

	// file open
	f, err := file.OpenFile(cLog.Filename)
	catchError(err)

	defer f.Close()

	// stat file
	stat, err := file.StatFile(f)
	catchError(err)

	cLog.Name = stat.Name()
	cLog.Size = stat.Size()

	// check initial date range of logs
	err = checkDateLog(f, cLog.LogStartTime, cLog.Size-1)
	catchError(err)

	// check final date range of logs
	err = checkDateLog(f, cLog.LogEndTime, 0)
	catchError(err)

	// graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		cancel()
	}()

	//
	file2.ProcessFile(ctx, f, cLog.LogStartTime, cLog.LogEndTime)
}

func init() {
	startTimeExec = time.Now()
	log.Logger.Info("Started")
}

func end() {
	diff := time.Since(startTimeExec)

	log.Logger.Info("\r")
	log.Logger.Info("Ended")
	log.Logger.Warning("Elapsed Time: ", diff.Truncate(time.Second).String())
}
