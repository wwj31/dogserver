protoc -I=%GOPATH%/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf -I=.\proto --gogofaster_out=.\ .\proto\*.proto

pause