package config

import "time"

type LogFile struct {
	Filename     string
	Name         string
	Size         int64
	LogStartTime *time.Time
	LogEndTime   *time.Time
}

type UserAgent struct {
	IsMobile       bool
	IsBoot         bool
	Model          string
	Plataform      string
	OSName         string
	OSVersion      string
	EngineName     string
	EngineVersion  string
	BrowserName    string
	BrowserVersion string
}

type LogRecord struct {
	Ip         string
	Date       *time.Time
	Method     string
	Resourse   string
	CodeStatus int
	Size       uint64
	UserAgent  *UserAgent
}
