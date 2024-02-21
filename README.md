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


## Develop Process
### Add a new channel
1. Select a channel, it can be websocket event or restful response
2. Base on the channel, add gRPC proto
3. Add the proto to genproto and generate protocol by makefile
4. Create service, implement according to pb in service.go
5. Register to cmd/server/main.go
6. Create a new client in cmd/client
7. Add build command in Makefile
8. Add new collection in script
9. Run and test
10. Add new Dockerfile in Dockerfiles/
11. Go to ECR to create a new registry
12. Go to Production

## Production
The app is running on AWS Fargate

For deployment, the pipeline is

- Push container to ECR

- Create task definition of ECS

- Deploy on ECS Cluster using Fargate