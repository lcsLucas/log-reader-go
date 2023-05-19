package regex

import (
	"errors"
	"log-reader-go/internal/config"
	"regexp"
	"strconv"
	"time"

	"github.com/mssola/useragent"
)

func LogParse(s string) (*config.LogRecord, error) {
	pattern, err := regexp.Compile(`(?P<IP>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s-\s-\s\[(?P<datahora>[^\[\]]+)\]\s"((?P<metodo>[A-Z]+)?\s(?P<recurso>[^"]+)?\sHTTP\/1\.[01]|.*)"\s(?P<status>\d{3})\s(?P<tamanho>\d+)\s"(?P<referer>[^"]+)?"\s"(?P<user_agent>[^"]+)?"`)

	if err != nil {
		return nil, err
	}

	match := pattern.FindStringSubmatch(s)

	if match == nil {
		return nil, errors.New("log record does not match")
	}

	l := &config.LogRecord{
		Ip:         match[1],
		Method:     match[4],
		Resourse:   match[5],
		Date:       nil,
		CodeStatus: 0,
		Size:       0,
	}

	t, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[2])
	if err != nil {
		return nil, err
	}
	l.Date = &t

	c, err := strconv.Atoi(match[6])
	if err != nil {
		return nil, err
	}
	l.CodeStatus = c

	sz, err := strconv.ParseUint(match[7], 10, 64)
	if err != nil {
		return nil, err
	}
	l.Size = sz

	ua := useragent.New(match[9])
	if ua != nil {
		enN, enV := ua.Engine()
		brN, brV := ua.Browser()

		l.UserAgent = &config.UserAgent{
			IsMobile:       ua.Mobile(),
			IsBoot:         ua.Bot(),
			Model:          ua.Model(),
			Plataform:      ua.Platform(),
			OSName:         ua.OSInfo().Name,
			OSVersion:      ua.OSInfo().Version,
			EngineName:     enN,
			EngineVersion:  enV,
			BrowserName:    brN,
			BrowserVersion: brV,
		}
	}

	return l, nil
}
