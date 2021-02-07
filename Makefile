.PHONY: build all release run run-client run-client-local run-server proto fmt clean

build:
	# native build
	go build -o tshooter *.go
all:
	# Linux
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o tshooter_linux *.go
	# Mac
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o tshooter_darwin *.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o tshooter_windows.exe *.go
release:
	cp assets/README.txt bin/
	cd bin && \
	zip tshooter_windows README.txt *_windows_* && \
	zip tshooter_linux README.txt *_linux_* && \
	zip tshooter_darwin README.txt *_darwin_*
run-client-local:
run:
	go run cmd/client_local.go
run-client:
	go run cmd/client.go
run-bot-client:
	go run cmd/bot_client.go
run-server:
	go run cmd/server.go
proto:
	protoc --go_out=plugins=grpc:. proto/*.proto
fmt:
	gofmt -s -w cmd/*.go proto/*.go pkg/*/*.go
clean:
	rm -f 