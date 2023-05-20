import 'package:grpc/grpc.dart';

import '../yock.pb.dart';
import 'yock.dart';

@Call("SayHello")
Future<CallResponse> doSomething(ServiceCall call, CallRequest request) async {
  print(request);
  return new CallResponse(buf: "I'm Dart");
}

void main(List<String> args) {
  YockInterface(args).Start();
}