@echo off
bash -c "%* >/mnt/c/Temp/_tmpout 2>/mnt/c/Temp/_tmperr"
if exist c:\Temp\_tmpout (
  type c:\Temp\_tmpout
  del c:\Temp\_tmpout
)
if exist c:\Temp\_tmperr (
  type c:\Temp\_tmperr 1>&2
  del c:\Temp\_tmperr
)
