protoc -I=.-I=%GOPATH%/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto -I=%GOPATH%/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf --gogofaster_out=./ -I=..\exec\  .\proto\*.proto
go generate ./inner/type.go
pause