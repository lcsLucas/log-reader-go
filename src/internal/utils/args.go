package utils

import (
	"errors"
	"flag"
	"log-reader-go/internal/config"
	"time"
)

func ReadArgs(cLog *config.LogFile) error {
	var err error

	filenameArg := flag.String("path", "", "File path")
	strLogStartTimeArg := flag.String("start", "", "Start of log period")
	strLogEndTimeArg := flag.String("end", "", "End of log period")

	flag.Parse()

	if len(*filenameArg) <= 0 {
		return errors.New("parameter \"path\" missing")
	}

	cLog.Filename = *filenameArg

	if len(*strLogStartTimeArg) > 0 {
		var logStartTime time.Time

		logStartTime, err = time.Parse("2006-01-02T15:04:05", *strLogStartTimeArg)

		if err != nil {
			logStartTime, err = time.Parse("2006-01-02", *strLogStartTimeArg)

			if err != nil {
				return errors.New("arg \"Start Time\"")
			}
		}

		cLog.LogStartTime = &logStartTime

	}

	if len(*strLogEndTimeArg) > 0 {
		var logEndTime time.Time
		logEndTime, err = time.Parse("2006-01-02T15:04:05", *strLogEndTimeArg)

		if err != nil {
			logEndTime, err = time.Parse("2006-01-02", *strLogEndTimeArg)

			if err != nil {
				return errors.New("arg \"End Time\"")
			}
		}

		cLog.LogEndTime = &logEndTime

	}

	return nil
}
