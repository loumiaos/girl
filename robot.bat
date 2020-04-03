SET GOPATH=%~dp0;%GOPATH%
del /f /s bin\robot.exe
go build -o bin/robot.exe robot
start /d %~dp0 bin/robot.exe

pause