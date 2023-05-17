# Log Reader Go

A simple service implemented in Golang to process large log files (Apache), storing them in an ElasticSearch database and visualizing them in Grafana.

# Uso

To build or run the `go run`, the following parameters must be provided:

`-path (Required):`
File path argument

`-start (Optional):`
Log start date argument (2006-01-02 OR 2006-01-02T15:04:05)

`-end (Optional):`
Log end date argument (2006-01-02 OR 2006-01-02T15:04:05)

Usage example:

- `go build -o cmd/main cmd/main.go`

- `./cmd/main -path "/var/log/log-reader-go/access.log" -start "2022-01-01" -end "2022-21-31"`
