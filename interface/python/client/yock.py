import grpc
from yocki.yocki_pb2_grpc import YockInterfaceStub
from yocki.yocki_pb2 import CallRequest, CallResponse


def YockDialCall(port: int, req: CallRequest) -> CallResponse:
    with grpc.insecure_channel('localhost:{port}'.format(port=port)) as channel:
        stub = YockInterfaceStub(channel)
        return stub.Call(req)
