protoc -I=. -I=%GOPATH%/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf --gogofaster_out=../inner_message/ -I=..\exec\  .\proto\*.proto
pause