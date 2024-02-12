# binance-mongo

## Initialize MongoDB
Check MongoUrl in mongo/initMongo.go
```
make initdb
```

## Generate protobuffer
```
make genproto
```

## Build
```
make build
```

## Run server
```
make server-run
```

## Run client
```
make wsdepth-run
```
client-run default pushes BTCUSDT to server.  In production, user can use 
```
./bin/$(BINARY)-client -logtostderr=true -v=2 -symbol=ETHUSDT
```
to specify to desire symbol

## Schema
There are three type conversion map to three different stage of the workflow
- binance.WsDepthEvent
binance.WsDepthEvent is the primitive data type defined by binance-go package
- pb.WsDepthEvent
pd.WsDepthEvent is the gRPC protocol, should be a 1-1 mapping from binance.WsDepthEvent
- model.WsDepthEvent
model.WsDepthEvent is the custom type to control the mongoDB's field, the content should be related to mongo/initdb.go

## Production
Run server
```
nohup ./bin/bmgo-server-linux-x86 -log_dir="./logs/" -stderrthreshold=INFO -vmodule=file=2 >/dev/null 2>&1 &
```

Run client
```
nohup ./bin/bmgo-wsdepth-linux-x86 -log_dir="./logs/" -stderrthreshold=INFO -vmodule=file=2 -symbol=BTCUSDT >/dev/null 2>&1 &
nohup ./bin/bmgo-rstdepth-linux-x86 -log_dir="./logs/" -stderrthreshold=INFO -vmodule=file=2 -symbol=BTCUSDT >/dev/null 2>&1 &
```