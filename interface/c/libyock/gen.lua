job("build", function(cenv)
    sh({
            redirect = true,
            debug = true
        }, "gcc -c ./yock.c -o libyock.o",
        "ar rcs libyock.a libyock.o",
        "go build -o libyock.dll -buildmode=c-shared",
        "go build -o libyock.a -buildmode=c-archive")
    return true
end)

job("clean_all", function(cenv)
    rm({
            redirect = true,
            debug = true
        }, "libyock.o", "cJSON.c", "cJSON.h",
        "libyock.dll", "libyock.a",
        "yock.pb.go", "yock_grpc.pb.go")
    return true
end)

job("clean", function(cenv)
    rm({
        redirect = true,
        debug = true
    }, "libyock.o", "libyock.dll", "libyock.a")
    return true
end)

jobs("all", "build", "clean")
