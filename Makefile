run:
	@go run ./cmd/main.go

build:
	@go build -o ./bin/gme_bin_app ./cmd/main.go

port:
	ps -ef | grep gme_bin_app

daemon:
	@nohup ./bin/gme_bin_app > server_log.txt &
