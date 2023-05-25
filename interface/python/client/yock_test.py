from client.yock import YockDialCall
from yocki.yock_pb2 import CallRequest

print(YockDialCall(9090, CallRequest(Fn="SayHello")))
