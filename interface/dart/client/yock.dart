import 'package:grpc/grpc.dart';

import '../yock.pbgrpc.dart';

class YockClient {
  late final ClientChannel channel;
  late final YockInterfaceClient cli;

  YockClient(int port) {
    this.channel = ClientChannel('localhost',
        port: port,
        options:
            const ChannelOptions(credentials: ChannelCredentials.insecure()));
    this.cli = YockInterfaceClient(channel);
  }

  Future<CallResponse> Call(CallRequest request) {
    return this.cli.call(request);
  }

  Future<void> close() async {
    this.channel.shutdown();
  }
}
