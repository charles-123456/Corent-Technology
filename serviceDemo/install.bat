@echo off
cd %0\..
serviceDemo.exe -service=install
if NOT "%errorlevel%" == "0" (
echo ************* serviceDemo Installation Failed *************
EXIT /B 1
) else (
echo ********** serviceDemo Installed Successfully *************
)
serviceDemo.exe -service=start
if NOT "%errorlevel%" == "0" (
echo ************* Starting serviceDemo  Failed *************
EXIT /B 1
) else (
echo ********** serviceDemo Started Successfully *************
)