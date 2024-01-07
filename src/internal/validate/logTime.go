package validate

import (
	"errors"
	"log-reader-go/internal/utils/regex"
	"time"
)

func ValidateDateRangeLogs(t *time.Time, lineRead []byte, after bool) error {
	reg, err := regex.LogParse(string(lineRead))

	if err != nil {
		return err
	}

	if after && reg.Date.After(*t) {
		return errors.New("log time cannot be earlier than start time")
	} else if !after && reg.Date.Before(*t) {
		return errors.New("log time cannot be earlier than start time")
	}

	return nil
}
