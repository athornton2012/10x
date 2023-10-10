## Run locally

`go mod tidy`

`go run main.go </path/to/csv>`

## Docker

`docker build . -t 10x`

`docker run -p 8080:8080 10x`

`curl -i 'http://localhost:8080/query?weather=rain'`

## Run tests

`ginkgo` or `go test ./...`