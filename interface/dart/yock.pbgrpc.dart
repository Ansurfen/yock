///
//  Generated code. Do not modify.
//  source: yock.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'yock.pb.dart' as $0;
export 'yock.pb.dart';

class YockInterfaceClient extends $grpc.Client {
  static final _$ping = $grpc.ClientMethod<$0.PingRequest, $0.PingResponse>(
      '/Yock.YockInterface/Ping',
      ($0.PingRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.PingResponse.fromBuffer(value));
  static final _$call = $grpc.ClientMethod<$0.CallRequest, $0.CallResponse>(
      '/Yock.YockInterface/Call',
      ($0.CallRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CallResponse.fromBuffer(value));
  static final _$info = $grpc.ClientMethod<$0.InfoRequest, $0.InfoResponse>(
      '/Yock.YockInterface/Info',
      ($0.InfoRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.InfoResponse.fromBuffer(value));

  YockInterfaceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.PingResponse> ping($0.PingRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$ping, request, options: options);
  }

  $grpc.ResponseFuture<$0.CallResponse> call($0.CallRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$call, request, options: options);
  }

  $grpc.ResponseFuture<$0.InfoResponse> info($0.InfoRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$info, request, options: options);
  }
}

abstract class YockInterfaceServiceBase extends $grpc.Service {
  $core.String get $name => 'Yock.YockInterface';

  YockInterfaceServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.PingRequest, $0.PingResponse>(
        'Ping',
        ping_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.PingRequest.fromBuffer(value),
        ($0.PingResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CallRequest, $0.CallResponse>(
        'Call',
        call_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CallRequest.fromBuffer(value),
        ($0.CallResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.InfoRequest, $0.InfoResponse>(
        'Info',
        info_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.InfoRequest.fromBuffer(value),
        ($0.InfoResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.PingResponse> ping_Pre(
      $grpc.ServiceCall call, $async.Future<$0.PingRequest> request) async {
    return ping(call, await request);
  }

  $async.Future<$0.CallResponse> call_Pre(
      $grpc.ServiceCall _call, $async.Future<$0.CallRequest> request) async {
    return call(_call, await request);
  }

  $async.Future<$0.InfoResponse> info_Pre(
      $grpc.ServiceCall call, $async.Future<$0.InfoRequest> request) async {
    return info(call, await request);
  }

  $async.Future<$0.PingResponse> ping(
      $grpc.ServiceCall call, $0.PingRequest request);
  $async.Future<$0.CallResponse> call(
      $grpc.ServiceCall call, $0.CallRequest request);
  $async.Future<$0.InfoResponse> info(
      $grpc.ServiceCall call, $0.InfoRequest request);
}
