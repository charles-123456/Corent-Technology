@echo off
cd %0\..
corent-go.exe -service=install
if NOT "%errorlevel%" == "0" (
echo ************* corent-go Installation Failed *************
EXIT /B 1
) else (
echo ********** corent-go Installed Successfully *************
)
corent-go.exe -service=start
if NOT "%errorlevel%" == "0" (
echo ************* Starting corent-go  Failed *************
EXIT /B 1
) else (
echo ********** corent-go Started Successfully *************
)