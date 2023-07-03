// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

#include <stdio.h>
#include "libyock/yocki.h"

CallRequest sayHello(CallRequest req)
{
    callRequest request = callRequestBuild(req);
    printf("Fn: %d, Arg: %d", request.fn, request.arg);
    return CallResponseBuild({.buf = "I'm C"});
}

int main(int argc, char **argv)
{
    YockBuilder();
    YockCall("SayHello", sayHello);
    YockRun();
    return 0;
}