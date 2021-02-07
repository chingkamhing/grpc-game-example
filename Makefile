.PHONY: build all release proto fmt clean

build:
	# native build
	go build -ldflags "-s -w" -o tshooter *.go
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
proto:
	protoc --go_out=plugins=grpc:. proto/*.proto
fmt:
	gofmt -s -w cmd/*.go proto/*.go pkg/*/*.go
clean:
	rm -f 