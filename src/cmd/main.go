package main

import (
	"context"
	"errors"
	"log-reader-go/internal/config"
	"log-reader-go/internal/utils/args"
	"log-reader-go/internal/utils/file"
	"log-reader-go/internal/utils/regex"
	file2 "log-reader-go/pkg/file"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	startTimeExec time.Time
	cLog          config.LogFile
)

func main() {
	defer end()

	/** Lendo os argumentos passados */
	err := args.Read(&cLog)

	if err != nil {
		log.Error(err.Error())
		return
	}

	/** Abrindo o arquivo de log */
	f, err := os.Open(cLog.Filename)

	if err != nil {
		log.Error(err.Error())
		return
	}

	defer f.Close()

	filestat, err := f.Stat()

	if err != nil {
		log.Error(err.Error())
		return
	}

	cLog.Name = filestat.Name()
	cLog.Size = filestat.Size()

	/** Validando a data ínicio com os registros do log */
	if cLog.LogStartTime != nil {

		offset := cLog.Size - 1

		lastLine, err := file.ReadLine(f, uint64(offset), true)

		if err != nil {
			log.Error(err.Error())
			return
		}

		reg, err := regex.LogParse(string(lastLine))

		if err != nil {
			log.Error(err.Error())
			return
		}

		if reg.Date.Before(*cLog.LogStartTime) {
			log.Error(errors.New("log time cannot be earlier than start time"))
			return
		}

	}

	/** Validando a data final com os registros do log */
	if cLog.LogEndTime != nil {

		offset := 0

		firstLine, err := file.ReadLine(f, uint64(offset), false)

		if err != nil {
			log.Error(err.Error())
			return
		}

		reg, err := regex.LogParse(string(firstLine))

		if err != nil {
			log.Error(err.Error())
			return
		}

		if reg.Date.After(*cLog.LogEndTime) {
			log.Error(errors.New("log time cannot be earlier than start time"))
			return
		}

	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		cancel()
	}()

	file2.ProcessFile(ctx, f, cLog.LogStartTime, cLog.LogEndTime)
}

func init() {
	startTimeExec = time.Now()
	log.Info("Started")
}

func end() {
	diff := time.Since(startTimeExec)

	log.Info("Ended")
	log.Warning("Elapsed Time: ", diff.Truncate(time.Second).String())
}
