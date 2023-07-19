from client.yock import YockDialCall
from yocki.yocki_pb2 import CallRequest

print(YockDialCall(8080, CallRequest(Fn="SayHello")))
