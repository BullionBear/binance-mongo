BINARY := bmgo

PACKAGE="github.com/BullionBear/binance-mongo"
VERSION := $(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null)
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP := $(shell date '+%Y-%m-%dT%H:%M:%S')
LDFLAGS := -X '${PACKAGE}/env.Version=${VERSION}' \
           -X '${PACKAGE}/env.CommitHash=${COMMIT_HASH}' \
           -X '${PACKAGE}/env.BuildTime=${BUILD_TIMESTAMP}'

initdb:
	go run ./scripts/initdb.go -logtostderr=true -v=2

genproto:
	protoc --go_out=. --go-grpc_out=. protocols/wsdepth.proto
	protoc --go_out=. --go-grpc_out=. protocols/rstdepth.proto
	protoc --go_out=. --go-grpc_out=. protocols/wspdepth.proto
	protoc --go_out=. --go-grpc_out=. protocols/wsmkstat.proto

build: server-build rstdepth-build wsdepth-build wspdepth-build

server-build:
	go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-server cmd/server/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-server-linux-x86 cmd/server/*.go

rstdepth-build:
	go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-rstdepth cmd/client/rstdepth/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-rstdepth-linux-x86 cmd/client/rstdepth/*.go

wsdepth-build:
	go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-wsdepth cmd/client/wsdepth/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-wsdepth-linux-x86 cmd/client/wsdepth/*.go

wspdepth-build:
	go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-wspdepth cmd/client/wspdepth/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-wspdepth-linux-x86 cmd/client/wspdepth/*.go

server-run:
	./bin/$(BINARY)-server -logtostderr=true -v=2

wsdepth-run:
	./bin/$(BINARY)-wsdepth -logtostderr=true -v=2 -symbol=BTCUSDT

wspdepth-run:
	./bin/$(BINARY)-wspdepth -logtostderr=true -v=2 -symbol=BTCUSDT

rstdepth-run:
	./bin/$(BINARY)-rstdepth -logtostderr=true -v=2 -symbol=BTCUSDT

clean:
	rm -rf bin/*
	rm -rf logs/*

# .PHONY: clean, build, run