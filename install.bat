@echo off

REM :START
REM set /p y="Are you planning to use normal mode instead of ssh-agent? [y/n] : "
REM 
REM if "%y%"=="y" (
REM   set opt=normal.go
REM ) else (
REM   if "%y%"=="n" (
REM     set opt=ssh-agent.go
REM   ) else (
REM     GOTO START
REM   )
REM )
REM 
REM go build gbash.go %opt%

go build gbash.go

REM set bbash=%windir%\system32\bbash.bat 
REM if exist %bbash% (
REM   DEL %bbash%
REM )

move gbash.exe %windir%\system32\
REM copy gbash.vbs %windir%\system32\
copy batch\*.bat %windir%\system32\
set /p y="Shall I install basic tools on bash for you? [y/n] : "
if "%y%"=="y" (
  bbash sudo apt install git curl wget gcc-mingw-w64-x86-64 -y
)
echo GttBoW Installation succeed.
