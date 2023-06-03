go run github.com/wwj31/msgidgen@v0.0.7 -pack=outer -path=./proto -tag=./outer -upper=false -prefix=Id

protoc -I=./proto --plugin protoc-gen-go=../exec/protoc-gen-go --go_out=./ ./proto/*.proto

go generate ./outer/type.go