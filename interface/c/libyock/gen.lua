--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

job("build", function(cenv)
    if is_exist("cJSON.h") then
        os.rename("cJSON.h", "cJSON.h.txt")
    end
    local repo = "https://raw.githubusercontent.com/DaveGamble/cJSON/master/"
    local libs = { "cJSON.c", "cJSON.h" }
    for _, lib in ipairs(libs) do
        curl({
            save = true,
            filename = function(s)
                return lib
            end,
            debug = true
        }, repo .. lib)
    end
    sh({
            redirect = true,
            debug = true
        }, "gcc -c ./yock.c -o libyock.o",
        "ar rcs libyock.a libyock.o",
        "go build -o libyock.dll -buildmode=c-shared",
        "go build -o libyock.a -buildmode=c-archive")
end)

job("clean_all", function(cenv)
    rm({
            redirect = true,
            debug = true
        }, "libyock.o", "cJSON.c", "cJSON.h",
        "libyock.dll", "libyock.a")
end)

job("clean", function(cenv)
    rm({
        redirect = true,
        debug = true
    }, "libyock.o", "libyock.dll", "libyock.a")
end)

job("recover", function(cenv)
    if is_exist("cJSON.h.txt") then
        os.rename("cJSON.h.txt", "cJSON.h")
    else
        local out, err = cat("cJSON.tpl")
        yassert(err)
        write_file("cJSON.h", out)
    end
end)

jobs("all", "build", "clean")
jobs("cr", "clean_all", "recover")
