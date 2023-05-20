import '../yock.pb.dart';
import 'yock.dart';

Future<void> main(List<String> args) async {
  var client = YockClient(9090);
  print(await client.Call(CallRequest()..fn = "SayHello"));
  await client.close();
}
