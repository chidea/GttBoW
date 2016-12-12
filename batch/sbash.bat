@echo off

SET _ARGS=%*
SET _ARGS=%_ARGS:C:\=/mnt/c/%
SET _ARGS=%_ARGS:\=/%

echo | set /p=%_ARGS% > c:\Temp\_tmpin
echo | set /p=" 2>/mnt/c/Temp/_tmperr" >> c:\Temp\_tmpin

FOR /F "tokens=* USEBACKQ" %%F IN (`isconemu`) DO (
  SET var=%%F
)
IF "%var%" NEQ "" (
  bash -cur_console:p /mnt/c/Temp/_tmpin
) ELSE (
  bash /mnt/c/Temp/_tmpin
)
del c:\Temp\_tmpin

if exist c:\Temp\_tmpout (
  type c:\Temp\_tmpout
  del c:\Temp\_tmpout
)
if exist c:\Temp\_tmperr (
  type c:\Temp\_tmperr 1>&2
  del c:\Temp\_tmperr
)
