package args

import (
	"errors"
	"flag"
	"log-reader-go/internal/config"
	"os"
	"time"
)

func Read(cLog *config.LogFile) error {
	var err error

	filenameArg := flag.String("path", "", "File path argument")
	strLogStartTimeArg := flag.String("start", "", "Log start date argument (2006-01-02 OR 2006-01-02T15:04:05)")
	strLogEndTimeArg := flag.String("end", "", "Log end date (2006-01-02 OR 2006-01-02T15:04:05) argument")

	flag.Parse()

	if len(*filenameArg) <= 0 {
		envFile := os.Getenv("ENV_FILE")

		if len(envFile) > 0 {
			*filenameArg = "/var/log/" + os.Getenv("ENV_PATH") + "/" + envFile
		} else {
			return errors.New("give command line \"path\" argument")
		}

	}

	cLog.Filename = *filenameArg

	if len(*strLogStartTimeArg) <= 0 {
		envStartTime := os.Getenv("ENV_STARTTIME")

		if len(envStartTime) > 0 {
			*strLogStartTimeArg = envStartTime
		}

	}

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

	if len(*strLogEndTimeArg) <= 0 {
		envEndTime := os.Getenv("ENV_ENDTIME")

		if len(envEndTime) > 0 {
			*strLogEndTimeArg = envEndTime
		}

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

	if cLog.LogStartTime != nil && cLog.LogEndTime != nil && cLog.LogEndTime.Before(*cLog.LogStartTime) {
		return errors.New("end time cannot be earlier than start time")
	}

	return nil
}
