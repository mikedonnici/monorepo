# go

Contains Go source code.

Todo:

- [x] Create basic grpc server
- [ ] Create python grpc client from same `proto` definition
- [ ] Create http service with fastapi that used grpc client to talk to gRPC server
- [ ] Add tests
- [ ] Set up a docker stack
- [ ] Store secrets / env vars in consul
- [ ] Deploy to GCS using terraform with GitHub actions

# grpcurl command line tool

- Install [grpcurl](https://github.com/fullstorydev/grpcurl)
- [Enable server reflection](https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md#enable-server-reflection) 

To see services (required reflection enabled):

```shell
grpcurl -plaintext 127.0.0.1:50051 list
```

To invoke an RPC:

```shell
grpcurl -d '{"name": "Leo"}' -plaintext 127.0.0.1:50051 hellopb.HelloService/SayHello
```

OR

```shell
grpcurl -d @ -plaintext 127.0.0.1:50051 hellopb.HelloService/SayHello <<EOM
{"name": "Leo"}
EOM
```

