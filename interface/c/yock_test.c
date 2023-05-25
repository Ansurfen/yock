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