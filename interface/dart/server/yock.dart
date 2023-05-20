import 'dart:io';
import 'dart:mirrors';

import 'package:grpc/grpc.dart';
import 'package:args/args.dart';
import '../yock.pbgrpc.dart';

typedef YockCall = Future<CallResponse> Function(ServiceCall, CallRequest);

class YockInterfaceService extends YockInterfaceServiceBase {
  late Map<String, YockCall> dict;

  YockInterfaceService() {
    this.dict = new Map();
  }

  @override
  Future<CallResponse> call(ServiceCall call, CallRequest request) async {
    if (this.dict.containsKey(request.fn)) {
      return this.dict[request.fn]!(call, request);
    }
    return CallResponse(buf: "unimplement method or bad request");
  }

  @override
  Future<InfoResponse> info(ServiceCall call, InfoRequest request) async {
    return InfoResponse();
  }

  @override
  Future<PingResponse> ping(ServiceCall call, PingRequest request) async {
    return PingResponse();
  }

  void Register(String name, YockCall call) {
    this.dict[name] = call;
  }

  void Unregister(String name) {
    this.dict.remove(name);
  }
}

class Call {
  final String name;
  const Call(this.name);
}

class YockInterface {
  late final Server server;
  late int port;

  YockInterface(List<String> args) {
    YockInterfaceService yock = new YockInterfaceService();
    final parser = ArgParser();
    parser.addOption('port', abbr: 'p');
    ArgResults result = parser.parse(args);
    try {
      port = int.parse(result['port']);
    } catch (e) {
      print("invalid port");
      exit(1);
    }
    currentMirrorSystem().libraries.forEach((_, lib) {
      lib.declarations.forEach((s, decl) {
        decl.metadata.where((m) => m.reflectee is Call).forEach((m) {
          var anno = m.reflectee as Call;
          if (decl is MethodMirror) {
            yock.Register(anno.name, (ServiceCall call, CallRequest req) async {
              return ((decl).owner as LibraryMirror)
                  .invoke(s, [call, req]).reflectee;
            });
          }
          ;
        });
      });
    });
    server = Server([yock]);
  }

  Future<void> Start() async {
    print("service start...");
    await this.server.serve(port: this.port);
  }
}
