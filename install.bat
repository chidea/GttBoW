@echo off

go build gbash.go ssh-agent.go
move gbash.exe %windir%\system32\
copy batch\*.bat %windir%\system32\

echo GttBoW Installation succeed.
