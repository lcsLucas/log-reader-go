package config

import "time"

type LogFile struct {
	Filename     string
	Name         string
	Size         int64
	LogStartTime *time.Time
	LogEndTime   *time.Time
}
