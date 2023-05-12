package config

import "time"

type LogFile struct {
	Filename     string
	LogStartTime *time.Time
	LogEndTime   *time.Time
}
