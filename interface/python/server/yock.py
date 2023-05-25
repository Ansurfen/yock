import argparse
import time
from concurrent import futures
from typing import Callable
import grpc
from yocki.yock_pb2_grpc import YockInterface, add_YockInterfaceServicer_to_server
from yocki.yock_pb2 import CallRequest, CallResponse

parser = argparse.ArgumentParser(
    description='Start a grpc server on the specified port.')
parser.add_argument('-p', '--port', metavar='port', type=int, default=0,
                    help='the port number to start the server on (default: 0)')
args = parser.parse_args()

if args.port == 0:
    raise "invalid port"


class YockInterfaceService(YockInterface):
    funcs: dict[str, Callable[[CallRequest], CallResponse]]

    def __init__(self) -> None:
        self.funcs = {}

    def Call(self, request: CallRequest, context) -> CallResponse:
        if request.Func in self.funcs:
            return self.funcs[request.Func](request)
        return CallResponse(Ok=False, Buf="unknown")

    def run(self) -> None:
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        add_YockInterfaceServicer_to_server(self, server)
        server.add_insecure_port('[::]:{port}'.format(port=args.port))
        server.start()
        print("start service...")
        try:
            while True:
                time.sleep(60 * 60 * 24)
        except KeyboardInterrupt:
            server.stop(0)

    def registerCallback(self, name, cb):
        self.funcs[name] = cb


Yocki = YockInterfaceService()


def Call(fn: str):
    def wrapper(func: Callable[[CallRequest], CallResponse]):
        Yocki.registerCallback(fn, func)
    return wrapper



