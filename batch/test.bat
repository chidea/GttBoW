@echo off
set ARG=%*
set ARG=%ARG:C:\=/mnt/c/%
set ARG=%ARG:\=/%
echo %ARG%