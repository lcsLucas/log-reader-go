package utils

import (
	"errors"
	"flag"
	"log-reader-go/internal/config"
	"time"
)

func ReadArgs(cLog *config.LogFile) error {
	var err error

	filenameArg := flag.String("path", "", "File path argument")
	strLogStartTimeArg := flag.String("start", "", "Log start date argument (2006-01-02 OR 2006-01-02T15:04:05)")
	strLogEndTimeArg := flag.String("end", "", "Log end date (2006-01-02 OR 2006-01-02T15:04:05) argument")

	flag.Parse()

	if len(*filenameArg) <= 0 {
		return errors.New("give command line \"path\" argument")
	}

	cLog.Filename = *filenameArg

	if len(*strLogStartTimeArg) > 0 {
		var logStartTime time.Time

		logStartTime, err = time.Parse("2006-01-02T15:04:05", *strLogStartTimeArg)

		if err != nil {
			logStartTime, err = time.Parse("2006-01-02", *strLogStartTimeArg)

			if err != nil {
				return errors.New("could not able to parse the start time")
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
				return errors.New("could not able to parse the finish time")
			}
		}

		cLog.LogEndTime = &logEndTime

	}

	return nil
}
