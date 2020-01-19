SET GOPATH=%~dp0;%GOPATH%
go build -o bin/robot.exe robot
start /d %~dp0 bin/robot.exe

pause