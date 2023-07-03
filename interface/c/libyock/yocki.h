// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

#include "libyock.h"

#define YockBuilder() Yock *hulo = newYock()
#define YockCall(name, callback) yockRegisterCall(hulo, name, callback)
#define YockRun() yockRun(hulo, argv[2])