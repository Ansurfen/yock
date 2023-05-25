exec({
    redirect = true,
    debug = true
}, "gcc ./yock_test.c ./libyock/cJSON.c -L ./libyock -lyock -o yock")