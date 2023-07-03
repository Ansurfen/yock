// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

#include "yock.h"
#include "cJSON.h"

String yockCall(Call cb, String req)
{
    return cb(req);
}

CallResponse callResponseBuild(callResponse res)
{
    cJSON *root = cJSON_CreateObject();
    cJSON_AddStringToObject(root, "Buf", res.buf);
    return cJSON_PrintUnformatted(root);
}

callRequest callRequestBuild(CallRequest req)
{
    cJSON *root = cJSON_Parse(req);
    if (!root)
        return (callRequest){.fn = "err", .arg = ""};
    callRequest ret;
    ret.fn = cJSON_GetStringValue(cJSON_GetObjectItem(root, "Fn"));
    ret.arg = cJSON_GetStringValue(cJSON_GetObjectItem(root, "Arg"));
    return ret;
}