@echo off

set command=%*

for /f "tokens=1,*" %%a in ("%command%") do (
    set cmd1=%%a
    set cmd2=%%b
)

powershell.exe -Command "Start-Process %cmd1% -Verb runAs -ArgumentList \"%cmd2%\""