import grpc

from fastapi import FastAPI, HTTPException

import hello_pb2
import hello_pb2_grpc

app = FastAPI()


@app.get("/hello")
async def hello(name: str = "you"):
    channel = grpc.insecure_channel('localhost:50051')
    stub = hello_pb2_grpc.HelloServiceStub(channel)
    req = hello_pb2.HelloRequest(name=name)
    try:
        res = stub.SayHello(req)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"{e}")
    return {"msg": res.text}
