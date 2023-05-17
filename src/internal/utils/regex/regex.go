package regex

import (
	"log-reader-go/internal/config"
	"regexp"
	"strconv"
	"time"

	"github.com/mssola/useragent"
)

func LogParse(s string) (config.LogRecord, error) {
	var l config.LogRecord
	pattern, err := regexp.Compile(`(?P<IP>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})\s-\s-\s\[(?P<datahora>[^\[\]]+)\]\s"(?P<metodo>[A-Z]+)\s(?P<recurso>[^"]+)\sHTTP\/1\.[01]"\s(?P<status>\d{3})\s(?P<tamanho>\d+)\s"(?P<referer>[^"]+)"\s"(?P<user_agent>[^"]+)"`)

	if err != nil {
		return l, err
	}

	matches := pattern.FindAllStringSubmatch(s, -1)

	for _, v := range matches {

		t, _ := time.Parse("02/Jan/2006:15:04:05 -0700", v[2])
		c, _ := strconv.Atoi(v[5])
		s, _ := strconv.ParseUint(v[6], 10, 64)

		l = config.LogRecord{
			Ip:         v[1],
			Date:       &t,
			Method:     v[3],
			Resourse:   v[4],
			CodeStatus: c,
			Size:       s,
		}

		ua := useragent.New(v[8])

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

	}

	return l, nil
}
