BINARY := bmgo

PACKAGE="github.com/BullionBear/binance-mongo"
VERSION := $(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null)
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP := $(shell date '+%Y-%m-%dT%H:%M:%S')
LDFLAGS := -X '${PACKAGE}/env.Version=${VERSION}' \
           -X '${PACKAGE}/env.CommitHash=${COMMIT_HASH}' \
           -X '${PACKAGE}/env.BuildTime=${BUILD_TIMESTAMP}'

build:
	go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-server cmd/server/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-server-linux-arm64 cmd/server/*.go
	go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-client cmd/client/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-client-linux-arm64 cmd/client/*.go

server-run:
	./bin/$(BINARY)-server -logtostderr=true

client-run:
	./bin/$(BINARY)-client -logtostderr=true

clean:
	rm -rf bin/*

# .PHONY: clean, build, run