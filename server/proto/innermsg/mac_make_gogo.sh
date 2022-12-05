export PATH=$PATH:../exec/

protoc -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf\
  -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto\
  -I=./proto\
  --gogofaster_out=../innermsg/ ./proto/*.proto

go generate ./inner/type.go

# $GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/extensions.md