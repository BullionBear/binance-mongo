BINARY := bmgo

PACKAGE="github.com/BullionBear/binance-mongo"
VERSION := $(shell git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null)
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_TIMESTAMP := $(shell date '+%Y-%m-%dT%H:%M:%S')
LDFLAGS := -X '${PACKAGE}/env.Version=${VERSION}' \
           -X '${PACKAGE}/env.CommitHash=${COMMIT_HASH}' \
           -X '${PACKAGE}/env.BuildTime=${BUILD_TIMESTAMP}'

initdb:
	go run ./script/initdb.go

genproto:
	protoc --go_out=. --go-grpc_out=. protocols/wsdepth.proto
	protoc --go_out=. --go-grpc_out=. protocols/rstdepth.proto

build:
	go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-server cmd/server/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-server-linux-x86 cmd/server/*.go
	go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-rstdepth cmd/client/rstdepth/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-rstdepth-linux-x86 cmd/client/rstdepth/*.go
	go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-wsdepth cmd/client/wsdepth/*.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o ./bin/$(BINARY)-wsdepth-linux-x86 cmd/client/wsdepth/*.go

server-run:
	./bin/$(BINARY)-server -logtostderr=true -v=2

wsdepth-run:
	./bin/$(BINARY)-wsdepth -logtostderr=true -v=2 -symbol=BTCUSDT

rstdepth-run:
	./bin/$(BINARY)-rstdepth -logtostderr=true -v=2 -symbol=BTCUSDT

clean:
	rm -rf bin/*

# .PHONY: clean, build, run