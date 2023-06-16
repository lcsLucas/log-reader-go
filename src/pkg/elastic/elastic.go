package elastic

import (
	"log-reader-go/internal/log"

	"github.com/elastic/go-elasticsearch/v8"
)

var es *elasticsearch.Client

func Eae() {

}

func init() {
	var err error

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://0.0.0.0:9200",
		},
	}

	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Logger.Panic(err)
	}

	res, err := es.Info()
	if err != nil {
		log.Logger.Panic(err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Logger.Panic(res.String())
	}

	log.Logger.Info(res)

}
