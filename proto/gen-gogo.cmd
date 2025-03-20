@echo on
protoc.exe -I. --gogogame_out=plugins=googleProto+newStringer:.. ./login/*.proto
@echo off
CALL :CHECK_FAIL

@echo on
protoc.exe -I. --gogogame_out=plugins=googleProto+newStringer:.. ./game/*.proto
@echo off
CALL :CHECK_FAIL

@echo on
gen-proto-error.exe --parseDir=./login,./game --txtPath=./error_code.txt --goPath=../errmsg --pkgName=errmsg
@echo off
CALL :CHECK_FAIL

:CHECK_FAIL
 @echo off
 if NOT ["%errorlevel%"]==["0"] (
     @echo ----------------------------- ---------------------------------------- ------------------------------------
     @echo ----------------------------- gen-proto-error failed,please check error ------------------------------------
     @echo ----------------------------- ---------------------------------------- ------------------------------------
     pause
     exit /b %errorlevel%
 )