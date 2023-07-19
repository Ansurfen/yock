from server.yock import Call, Yocki
from yocki.yocki_pb2 import CallRequest, CallResponse


@Call(fn="SayHello")
def sayHello(request: CallRequest) -> CallResponse:
    print("recv: {}".format(request))
    return CallResponse(Buf="I'm Python")


Yocki.run()
