export PATH=$PATH:../exec/

go run github.com/wwj31/msgidgen@latest -pack=outer -path=./proto -tag=./outer -upper=false -prefix=Id

protoc -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf\
 -I=./proto\
 --gogofaster_out=./ ./proto/*.proto

go generate ./outer/type.go