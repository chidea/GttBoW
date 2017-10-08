@echo off

taskkill /f /im gbash.exe
go build gbash.go

set bbash=%windir%\system32\bbash.bat 
if exist %bbash% (
  DEL %bbash%
  echo Batch file of old versions removed
)
move gbash.exe %windir%\system32\
echo bash proxy binary(gbash) installed

FOR %%G IN (git) DO go build -ldflags "-X main.cmd=%%G" -o %%G.exe exeproxy/proxy.go && move %%G.exe %windir%\system32\
echo git proxy binary installed

set /p y="Do you want ssh and scp binaries be installed? [y/n] : "
IF "%y%"=="y" (
  FOR %%G IN (scp, ssh) DO go build -ldflags "-X main.cmd=%%G" -o %%G.exe exeproxy/proxy.go && move %%G.exe %windir%\system32\
  echo ssh/scp proxy binary Installed
)

set /p y="Do you want curl and wget binary be installed? [y/n] : "
IF "%y%"=="y" (
  FOR %%G IN (curl, wget) DO go build -ldflags "-X main.cmd=%%G" -o %%G.exe exeproxy/proxy.go && move %%G.exe %windir%\system32\
)

set /p y="Do you want make and man binary be installed? [y/n] : "
IF "%y%"=="y"(
  FOR %%G IN (make, man) DO go build -ldflags "-X main.cmd=%%G" -o %%G.exe exeproxy/proxy.go && move %%G.exe %windir%\system32\
)

set /p y="Do you want gcc related binaries be installed? [y/n] : "
IF "%y%"=="y" (
  FOR %%G IN (addr2line, cpp, dllwrap, elfedit, gcc-ar, gcc-nm, gcc-ranlib, gcc, gcov, gprof, ld, nm, objcopy, objdump, ranlib, readelf, size, strings, strip, windmc, windres) DO go build -ldflags "-X main.cmd=x86_64-w64-mingw32-%%G" -o %%G.exe exeproxy/proxy.go && move %%G.exe %windir%\system32\
  echo GCC proxy binary Installed
)

set /p y="Shall I install basic tools like git and mingw on bash for you? [y/n] : "
IF "%y%"=="y" (
  gbash sudo apt install git curl wget gcc-mingw-w64-x86-64 -y
)

set /p y="Do you want ssh-agent installed on bash? [y/n] : "
IF "%y%"=="y" (
  powershell .\install-ssh-agent.ps1
  echo ssh-agent script installed
)

echo GttBoW Installation succeed.
