@echo off

go build gbash.go ssh-agent.go
move gbash.exe %windir%\system32\
copy batch\*.bat %windir%\system32\
bbash sudo apt install git curl wget gcc-mingw-w64-x86-64 -y

echo GttBoW Installation succeed.
