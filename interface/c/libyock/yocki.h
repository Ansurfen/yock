#include "libyock.h"

#define YockBuilder() Yock *hulo = newYock()
#define YockCall(name, callback) yockRegisterCall(hulo, name, callback)
#define YockRun() yockRun(hulo, argv[2])