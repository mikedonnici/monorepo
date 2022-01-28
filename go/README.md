# go

Contains Go source code.

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

