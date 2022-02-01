protoc -I=. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf --gogofaster_out=../innermsg/ -I=../exec/  ./proto/*.proto
go generate ./outer/type.go