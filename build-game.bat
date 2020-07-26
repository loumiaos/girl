SET GOPATH=%~dp0;%GOPATH%
cd src/game
go build -o ../../bin/game/game64.exe game

pause