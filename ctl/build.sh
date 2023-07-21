#!/bin/bash

SCRIPT_PATH=$(pwd)

ffi=0
dev=0
remote=0
os=linux

for arg in "$@"; do
str=$(echo $arg | cut -c1-2)
if [ $arg = ffi ]
then
    ffi=1
elif [ $arg = dev ]
then
    dev=1
elif [ $arg = remote ]
then
    remote=1
elif [  $str = os ]
then
    os=$(echo $arg | cut -c3-${#arg})
else
    echo "unknown flag"
fi                 
done

cd ../scheduler
if [ $ffi = 0 ]
then
	mv yockf.go yockf.go.txt
else
	mv yockf.go.txt yockf.go
fi

cd $SCRIPT_PATH

go env -w GOOS=linux

if [ $dev = 0 ]
then
	go run . run ../auto/build.lua all -- --all-os $os --all-r $remote
else
	go run . run ../auto/build.lua alldev -- --alldev-os $os --alldev-r $remote
fi

if [ $ffi = 1 ]
then
	cd ../scheduler
	mv yockf.go yockf.go.txt
	cd $SCRIPT_PATH
fi