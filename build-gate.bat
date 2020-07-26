SET GOPATH=%~dp0;%GOPATH%
cd src/gate
go build -o ../../bin/gate/gate64.exe gate

pause