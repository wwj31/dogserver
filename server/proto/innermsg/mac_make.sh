../exec/protoc --plugin protoc-gen-go=../exec/protoc-gen-go -I=./ --go_out=./ ./proto/*.proto
go generate ./inner/type.go