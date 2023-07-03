#!/bin/bash

if [ -z "$1" ]; then
    go run . run ../auto/build.lua all
elif [ "$1" == "dev" ]; then
    # dev environment
    go run . run ../auto/build.lua all-dev
else
    echo "invalid task"
fi
