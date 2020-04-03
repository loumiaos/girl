SET GOPATH=%~dp0;%GOPATH%
del /f /s bin\game\game64.exe
del /f /s bin\gate\gate64.exe
del /f /s bin\login\login64.exe

go build -o bin/game/game64.exe game
go build -o bin/gate/gate64.exe gate
go build -o bin/login/login64.exe login

start /d "bin\game" game64.exe
start /d "bin\login" login64.exe
start /d "bin\gate" gate64.exe

pause