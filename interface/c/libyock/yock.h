// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

#ifndef __YOCK_H__
#define __YOCK_H__

typedef char *String;
typedef String error;
typedef String (*Call)(String);
typedef char Boolean;
typedef String CallResponse;
typedef String CallRequest;
typedef struct
{
    void *ptr;
} Yock;

String yockCall(Call, String);

typedef struct
{
    String buf;
} callResponse;

typedef struct
{
    String fn;
    String arg;
} callRequest;

callRequest callRequestBuild(CallRequest);
CallResponse callResponseBuild(callResponse);

#define CallResponseBuild(req) callResponseBuild((callResponse)req)

#endif