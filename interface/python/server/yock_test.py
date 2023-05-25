from server.yock import Call, Yocki
from yocki.yock_pb2 import CallRequest, CallResponse


@Call(fn="SayHello")
def sayHello(request: CallRequest):
    print("recv: ", request)
    return CallResponse(Buf="I'm Python")

Yocki.run()