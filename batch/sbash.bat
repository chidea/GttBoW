@echo off

echo %* | set /p=%* > c:\Temp\_tmpin
echo ^>/mnt/c/Temp/_tmpout 2^>/mnt/c/Temp/_tmperr >> c:\Temp\_tmpin
bash /mnt/c/Temp/_tmpin
del c:\Temp\_tmpin

if exist c:\Temp\_tmpout (
  type c:\Temp\_tmpout
  del c:\Temp\_tmpout
)
if exist c:\Temp\_tmperr (
  type c:\Temp\_tmperr 1>&2
  del c:\Temp\_tmperr
)
