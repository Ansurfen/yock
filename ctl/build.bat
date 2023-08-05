@REM build.bat 
@REM build.bat dev
@REM build.bat dev ffi oswindows/oslinux

@echo off

set "script_path=%~dp0"
cd %script_path%

set "ffi=0"
set "dev=0"
set "os=windows"
set "ver=0"
set "remote=0"

for %%A in (%*) do (
    if /I "%%A" == "ffi" (
        set "ffi=1"
    ) else if /I "%%A" == "dev" (
        set "dev=1"
    ) else if /I "%%A" == "remote" (
        set "remote=1"
    ) else (
        setlocal enabledelayedexpansion
        set "prefix=%%A"
        if "!prefix:~0,2!" == "os" (
            set os=!prefix:~2!
        ) else if "!prefix:~0,1!" == "v" (
            set ver=!prefix:~1!
        ) else (
            echo unknown argument: %%A
        )
    )
)

cd ../scheduler
if %ffi% equ 0 (
    @REM echo comments ffi
    if exist "yockf.go" (
        ren yockf.go yockf.go.txt
    )
) else (
    if exist "yockf.go.txt" (
        ren yockf.go.txt yockf.go
    )
)
cd %script_path%

go env -w GOOS=windows

if %dev% equ 0 (
    go run . run ../auto/build.lua all -- --all-os %os% --all-r %remote% --all-v %ver%
) else (
    @REM dev environment
    go run . run ../auto/build.lua alldev -- --alldev-os %os% --alldev-r %remote% --alldev-v %ver%
)

if %ffi% equ 1 (
    @REM echo recovery ffi file
    cd ../scheduler
    if exist "yockf.go" (
        ren yockf.go yockf.go.txt
    )
    cd %script_path%
)
