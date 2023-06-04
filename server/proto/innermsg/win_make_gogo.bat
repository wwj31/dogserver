..\exec\protoc.exe -I=./proto -I=%GOPATH%\pkg\mod\github.com\gogo\protobuf@v1.3.2\gogoproto --gogofaster_out=./ -I=..\exec\  .\proto\*.proto
go generate ./inner/type.go
pause