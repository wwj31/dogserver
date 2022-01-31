protoc -I=./proto --plugin protoc-gen-go=../exec/protoc-gen-go --go_out=./ ./proto/*.proto
go generate ./outer/type.go