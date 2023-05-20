///
//  Generated code. Do not modify.
//  source: yock.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use pingRequestDescriptor instead')
const PingRequest$json = const {
  '1': 'PingRequest',
};

/// Descriptor for `PingRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pingRequestDescriptor = $convert.base64Decode('CgtQaW5nUmVxdWVzdA==');
@$core.Deprecated('Use pingResponseDescriptor instead')
const PingResponse$json = const {
  '1': 'PingResponse',
};

/// Descriptor for `PingResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List pingResponseDescriptor = $convert.base64Decode('CgxQaW5nUmVzcG9uc2U=');
@$core.Deprecated('Use callRequestDescriptor instead')
const CallRequest$json = const {
  '1': 'CallRequest',
  '2': const [
    const {'1': 'Fn', '3': 1, '4': 1, '5': 9, '10': 'Fn'},
    const {'1': 'Arg', '3': 2, '4': 1, '5': 9, '10': 'Arg'},
  ],
};

/// Descriptor for `CallRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List callRequestDescriptor = $convert.base64Decode('CgtDYWxsUmVxdWVzdBIOCgJGbhgBIAEoCVICRm4SEAoDQXJnGAIgASgJUgNBcmc=');
@$core.Deprecated('Use callResponseDescriptor instead')
const CallResponse$json = const {
  '1': 'CallResponse',
  '2': const [
    const {'1': 'Buf', '3': 1, '4': 1, '5': 9, '10': 'Buf'},
  ],
};

/// Descriptor for `CallResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List callResponseDescriptor = $convert.base64Decode('CgxDYWxsUmVzcG9uc2USEAoDQnVmGAEgASgJUgNCdWY=');
@$core.Deprecated('Use infoRequestDescriptor instead')
const InfoRequest$json = const {
  '1': 'InfoRequest',
};

/// Descriptor for `InfoRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List infoRequestDescriptor = $convert.base64Decode('CgtJbmZvUmVxdWVzdA==');
@$core.Deprecated('Use infoResponseDescriptor instead')
const InfoResponse$json = const {
  '1': 'InfoResponse',
};

/// Descriptor for `InfoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List infoResponseDescriptor = $convert.base64Decode('CgxJbmZvUmVzcG9uc2U=');
