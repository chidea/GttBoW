@echo off

REM SET _ARGS=%*
REM SET _ARGS=%_ARGS:C:\=/mnt/c/%
REM SET _ARGS=%_ARGS:\=/%

REM uncomment bellow to use ssh-agent
REM echo | set /p="if [[ $(pgrep ssh-agent) != '' ]]; then . ~/.ssh/ssh-agent.sh; fi;" > c:\Temp\_tmpin

echo | set /p=%* > c:\Temp\_tmpin
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
