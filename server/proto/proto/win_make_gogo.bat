protoc -I=%GOPATH%/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf -I=.\ --gogofaster_out=..\ .\*.proto

pause