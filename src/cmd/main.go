package main

import (
	"errors"
	"log-reader-go/internal/config"
	"log-reader-go/internal/utils/args"
	"log-reader-go/internal/utils/file"
	"log-reader-go/internal/utils/regex"
	"log-reader-go/pkg/color"
	file2 "log-reader-go/pkg/file"
	"os"
	"time"
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
		color.PrintError(err.Error())
		return
	}

	/** Abrindo o arquivo de log */
	f, err := os.Open(cLog.Filename)

	if err != nil {
		color.PrintError(err.Error())
		return
	}

	defer f.Close()

	filestat, err := f.Stat()

	if err != nil {
		color.PrintError(err.Error())
		return
	}

	cLog.Name = filestat.Name()
	cLog.Size = filestat.Size()

	/** Validando a data ínicio com os registros do log */
	if cLog.LogStartTime != nil {

		offset := cLog.Size - 1

		lastLine, err := file.ReadLine(f, uint64(offset), true)

		if err != nil {
			color.PrintError(err.Error())
			return
		}

		reg, err := regex.LogParse(string(lastLine))

		if err != nil {
			color.PrintError(err.Error())
			return
		}

		if reg.Date.Before(*cLog.LogStartTime) {
			color.PrintError(errors.New("log time cannot be earlier than start time"))
			return
		}

	}

	/** Validando a data final com os registros do log */
	if cLog.LogEndTime != nil {

		offset := 0

		firstLine, err := file.ReadLine(f, uint64(offset), false)

		if err != nil {
			color.PrintError(err.Error())
			return
		}

		reg, err := regex.LogParse(string(firstLine))

		if err != nil {
			color.PrintError(err.Error())
			return
		}

		if reg.Date.After(*cLog.LogEndTime) {
			color.PrintError(errors.New("log time cannot be earlier than start time"))
			return
		}

	}

	file2.ProcessFile(f, cLog.LogStartTime, cLog.LogEndTime)

}

func init() {
	startTimeExec = time.Now()
	color.PrintCyan("Started: ", startTimeExec.String())
}

func end() {
	diff := time.Since(startTimeExec)

	color.PrintCyan("Ended: ", time.Now().String())
	color.PrintYellow("Elapsed Time: ", diff.Truncate(time.Second).String())
}
